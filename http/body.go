package http

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"
)

func ExtractBody(response *http.Response) ([]byte, error) {
	if response == nil {
		return nil, fmt.Errorf("'http.Response' must not be a nil")
	}

	//goland:noinspection ALL
	defer response.Body.Close()

	if response.ContentLength != -1 && response.ContentLength < 1 {
		return []byte{}, nil
	}

	cEncoding := response.Header.Get(Header.ContentEncoding)
	cType := response.Header.Get(Header.ContentType)
	reader := response.Body

	if reader == nil {
		return []byte{}, nil
	}

	//goland:noinspection ALL
	if "gzip" == cEncoding {
		readerZ, erZ := gzip.NewReader(reader)
		defer reader.Close()

		if erZ == nil {
			reader = readerZ
		}
	} else if "deflate" == cEncoding {
		readerZ, erZ := zlib.NewReader(reader)
		defer reader.Close()

		if erZ == nil {
			reader = readerZ
		}
	}

	var (
		body []byte
		er   error
	)

	//goland:noinspection ALL
	if "application/octet-stream" != cType && len(strings.TrimSpace(cType)) == 0 {
		utf8, err := charset.NewReader(reader, cType)
		defer reader.Close()

		if err != nil {
			return nil, fmt.Errorf("could not decode response body; %s", err)
		}

		body, er = ioutil.ReadAll(utf8)
	} else {
		body, er = ioutil.ReadAll(reader)
	}

	if er != nil {
		return nil, fmt.Errorf("could not read response body")
	}

	return body, nil
}
