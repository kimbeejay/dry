package http

import "net/http"

type Request interface {
	GetHeaders() map[string]string
	GetMethod() string
	GetEndpoint() string
	GetCookies() []http.Cookie
	GetValues() string
	GetPayload() []byte
}
