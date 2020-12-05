package util

import (
	"bytes"
	"net/http"
	"strings"
)

func GetRequestAsString(r *http.Request) string {
	buf := new(bytes.Buffer)
	if err := r.Write(buf); err != nil {
		return err.Error()
	}
	return string(buf.Bytes())
}

func IsFromCurl(r *http.Request) bool {
	userAgent := r.Header.Get("User-Agent")
	return strings.HasPrefix(userAgent, "curl/")
}
