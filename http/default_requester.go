package http

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/golang/glog"
	dString "github.com/kimbeejay/dry/string"
)

type defaultRequester struct {
	ctx   context.Context
	host  string
	debug bool
}

//goland:noinspection GoUnusedExportedFunction
func NewDefaultRequester(ctx context.Context, host string, debug bool) *defaultRequester {
	return &defaultRequester{
		ctx:   ctx,
		host:  host,
		debug: debug,
	}
}

func (r *defaultRequester) SetHost(host string) {
	r.host = host
}

func (r *defaultRequester) Do(q Request, goodCodes []int) (int, []byte, error) {
	if q == nil {
		return 0, nil, fmt.Errorf("'IRequest' must not be a nil")
	}

	if len(goodCodes) < 1 {
		return 0, nil, fmt.Errorf("please, specify more good codes")
	}

	endpoint, er := r.makeEndpoint(q)
	if er != nil {
		return 0, nil, er
	}

	if allowed, ok := knownMethods[q.GetMethod()]; !ok || !allowed {
		return 0, nil, fmt.Errorf("method '%s' is unknown or not allowed for requesting", q.GetMethod())
	}

	var payload *bytes.Reader
	if q.GetPayload() != nil &&
		(http.MethodPost == q.GetMethod() || http.MethodPut == q.GetMethod()) {
		payload = bytes.NewReader(q.GetPayload())
	} else {
		payload = bytes.NewReader(nil)
	}

	request, er := http.NewRequestWithContext(r.ctx, q.GetMethod(), endpoint.String(), payload)
	if er != nil {
		r.DebugLog(fmt.Sprintf("could not build new 'http.Request': %v", er))
	}

	// Apply request headers;
	for k, v := range q.GetHeaders() {
		request.Header.Set(k, v)
	}

	// Apply request cookies;
	for _, c := range q.GetCookies() {
		request.AddCookie(&c)
	}

	client := ProduceDefaultClient()
	response, er := client.Do(request)
	if er != nil {
		return 0, nil, er
	}

	body, er := ExtractBody(response)
	if er != nil {
		return response.StatusCode, nil, er
	}

	isGood := false
	for i := range goodCodes {
		if response.StatusCode == goodCodes[i] {
			isGood = true
			break
		}
	}

	if isGood {
		return response.StatusCode, body, nil
	} else {
		return response.StatusCode, nil, fmt.Errorf("got unexpected http status code %d: %s", response.StatusCode, body)
	}
}

func (r *defaultRequester) DebugLog(s ...string) {
	if !r.debug {
		return
	}

	glog.Infoln(s)
}

func (r *defaultRequester) makeEndpoint(q Request) (*url.URL, error) {
	endpoint, er := url.Parse(fmt.Sprintf("%s/%s", r.host, q.GetEndpoint()))
	if er != nil {
		return nil, fmt.Errorf("could not make new endpoint: %v", er)
	}

	if !dString.IsValidUrl(endpoint.String()) {
		return nil, fmt.Errorf("could not make new endpoint: %v", er)
	}

	if !dString.IsEmpty(q.GetValues()) {
		endpoint.RawQuery = q.GetValues()
	}

	return endpoint, nil
}
