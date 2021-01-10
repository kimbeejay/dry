package http

type headerList struct {
	Accept          string
	AcceptEncoding  string
	Authorization   string
	ContentEncoding string
	ContentType     string
}

var Header = &headerList{
	Accept:          "Accept",
	AcceptEncoding:  "Accept-Encoding",
	Authorization:   "Authorization",
	ContentEncoding: "Content-Encoding",
	ContentType:     "Content-Type",
}

var DefaultHttpHeaders = map[string]string{
	Header.Accept:         "application/json",
	Header.AcceptEncoding: "gzip, deflate",
	Header.ContentType:    "application/json; charset=UTF-8",
}
