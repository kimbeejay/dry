package http

import (
	"net/http"
	"time"
)

const DefaultTimeout = time.Second * 30

func ProduceDefaultClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport)
	t.Proxy = http.ProxyFromEnvironment

	client := new(http.Client)
	client.Transport = t
	client.Timeout = DefaultTimeout

	return client
}
