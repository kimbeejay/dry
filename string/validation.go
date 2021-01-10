package string

import (
	"net/url"
	"strings"
)

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsValidUrl(s string) bool {
	l, er := url.Parse(s)
	return er == nil &&
		!IsEmpty(l.Scheme) &&
		!IsEmpty(l.Host)
}
