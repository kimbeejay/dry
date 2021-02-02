package http

import (
	"net/http"
	"strings"
)

var knownMethods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

func IsKnownMethod(s string) bool {
	for i := range knownMethods {
		if strings.Compare(knownMethods[i], s) == 0 {
			return true
		}
	}

	return false
}
