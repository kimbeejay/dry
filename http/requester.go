package http

import "net/url"

type Requester interface {
	Do(Request, []int) (Response, error)

	DebugLog(...string)

	GetEndpoint(Request) (*url.URL, error)
}
