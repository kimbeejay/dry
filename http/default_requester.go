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

const (
	defaultRequesterCookies = "DefaultRequester.Ctx.Cookies"
	defaultRequesterHost    = "DefaultRequester.Ctx.Host"
)

type defaultRequester struct {
	ctx   context.Context
	debug bool
}

//goland:noinspection GoUnusedExportedFunction
func NewDefaultRequester(ctx context.Context, host string, debug bool) *defaultRequester {
	requester := &defaultRequester{
		ctx:   context.WithValue(ctx, defaultRequesterHost, host),
		debug: debug,
	}

	return requester
}

func (r *defaultRequester) Do(q Request, goodCodes []int) (Response, error) {
	if q == nil {
		return nil, fmt.Errorf("'IRequest' must not be a nil")
	}

	if len(goodCodes) < 1 {
		return nil, fmt.Errorf("please, specify more good codes")
	}

	endpoint, er := r.makeEndpoint(q)
	if er != nil {
		return nil, er
	}

	if allowed, ok := knownMethods[q.GetMethod()]; !ok || !allowed {
		return nil, fmt.Errorf("method '%s' is unknown or not allowed for requesting", q.GetMethod())
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

	// Apply context cookies;
	for _, cookie := range r.getCookies() {
		request.AddCookie(&cookie)
	}

	// Apply request cookies;
	for _, c := range q.GetCookies() {
		request.AddCookie(&c)
	}

	client := ProduceDefaultClient()
	response, er := client.Do(request)
	if er != nil {
		return nil, er
	}

	res := newDefaultResponse(
		response.StatusCode,
		response.Header,
		ExtractCookies(response),
		nil)

	if body, er := ExtractBody(response); er != nil {
		return res, er
	} else {
		res.body = body
	}

	isGood := false
	for i := range goodCodes {
		if response.StatusCode == goodCodes[i] {
			isGood = true
			break
		}
	}

	if isGood {
		cookies := res.Cookies()
		if len(cookies) > 0 {
			r.ctx = context.WithValue(r.ctx, defaultRequesterCookies, cookies)
		}

		return res, nil
	} else {
		e := res.body
		res.body = nil
		return res, fmt.Errorf("got unexpected http status code %d: %s", response.StatusCode, e)
	}
}

func (r *defaultRequester) DebugLog(s ...string) {
	if !r.debug {
		return
	}

	glog.Infoln(s)
}

func (r *defaultRequester) makeEndpoint(q Request) (*url.URL, error) {
	absLink := q.GetEndpoint()
	if !ContainsKnownScheme(q.GetEndpoint()) {
		host := r.getHost()
		if dString.IsEmpty(host) {
			return nil, fmt.Errorf("could not makeEndpoint: host must not be an empty string")
		}

		absLink = fmt.Sprintf("%s/%s", host, q.GetEndpoint())
	}

	endpoint, er := url.Parse(absLink)
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

func (r *defaultRequester) getCookies() []http.Cookie {
	if c, ok := r.ctx.Value(defaultRequesterCookies).([]http.Cookie); ok {
		return c
	}

	return []http.Cookie{}
}

func (r *defaultRequester) getHost() string {
	if host, ok := r.ctx.Value(defaultRequesterHost).(string); ok {
		return host
	}

	return ""
}
