package http

import "net/http"

func ExtractCookies(response *http.Response) []http.Cookie {
	if response == nil {
		return []http.Cookie{}
	}

	cookies := make([]http.Cookie, len(response.Cookies()))
	for i, c := range response.Cookies() {
		cookies[i] = http.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    c.Expires,
			RawExpires: c.RawExpires,
			MaxAge:     c.MaxAge,
			Secure:     c.Secure,
			HttpOnly:   c.HttpOnly,
			SameSite:   c.SameSite,
			Raw:        c.Raw,
			Unparsed:   c.Unparsed,
		}
	}

	return cookies
}
