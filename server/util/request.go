package util

import (
	"bytes"
	"net/http"
	"net/http/httputil"
	"strings"
)

func GetRequestAsString(r *http.Request) string {
	requestBytes, err := httputil.DumpRequest(r, false)
	if err != nil {
		return err.Error()
	}
	return string(bytes.TrimSpace(requestBytes))
}

func IsFromCurl(r *http.Request) bool {
	userAgent := r.Header.Get("User-Agent")
	return strings.HasPrefix(userAgent, "curl/")
}
