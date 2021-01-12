package http

type headerList struct {
	Accept          string
	AcceptEncoding  string
	AcceptLanguage  string
	Authorization   string
	ContentEncoding string
	ContentType     string
	UserAgent       string
}

var Header = &headerList{
	Accept:          "Accept",
	AcceptEncoding:  "Accept-Encoding",
	AcceptLanguage:  "Accept-Language",
	Authorization:   "Authorization",
	ContentEncoding: "Content-Encoding",
	ContentType:     "Content-Type",
	UserAgent:       "User-Agent",
}

var DefaultHttpHeaders = map[string]string{
	Header.Accept:         "application/json",
	Header.AcceptEncoding: "gzip, deflate",
	Header.ContentType:    "application/json; charset=UTF-8",
}
