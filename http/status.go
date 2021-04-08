package http

type Status int

const (
	Information Status = iota
	Successful
	Redirection
	ClientError
	ServerError
)

func StatusBelongs(status Status, code int) bool {
	switch status {
	case Information: // [100-199]
		return code > 99 && code < 200
	case Successful:  // [200-299]
		return code > 199 && code < 300
	case Redirection: // [300-399]
		return code > 299 && code < 400
	case ClientError: // [400-499]
		return code > 399 && code < 500
	case ServerError: // [500-599]
		return code > 499 && code < 600
	}

	return false
}
