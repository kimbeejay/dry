package http

import (
	"net/http"
)

type defaultResponse struct {
	statusCode int
	headers    http.Header
	cookies    []http.Cookie
	body       []byte
}

func NewDefaultResponse(code int, h http.Header, c []http.Cookie, b []byte) *defaultResponse {
	r := new(defaultResponse)
	r.statusCode = code
	r.headers = h
	r.cookies = c
	r.body = b
	return r
}

func (d *defaultResponse) StatusCode() int {
	return d.statusCode
}

func (d *defaultResponse) Headers() http.Header {
	return d.headers
}

func (d *defaultResponse) Cookies() []http.Cookie {
	return d.cookies
}

func (d *defaultResponse) Body() []byte {
	return d.body
}
