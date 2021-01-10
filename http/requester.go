package http

import "net/url"

type Requester interface {
	Do(Request, []int) (int, []byte, error)

	DebugLog(...string)

	makeEndpoint(Request) (*url.URL, error)
}
