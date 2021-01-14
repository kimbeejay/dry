package http

import "net/http"

func ExtractCookies(response *http.Response) []http.Cookie {
	cookies := make([]http.Cookie, 0)
	if response == nil {
		return cookies
	}

	for _, c := range response.Cookies() {
		cookies = append(cookies, http.Cookie{
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
		})
	}

	return cookies
}
