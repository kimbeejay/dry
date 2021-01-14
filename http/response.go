package http

import "net/http"

type Response interface {
	StatusCode() int
	Headers() http.Header
	Cookies() []http.Cookie
	Body() []byte
}
