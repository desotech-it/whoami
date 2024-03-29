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

func MakeClientInfo(r *http.Request) map[string]string {
	return map[string]string{
		"client_address": strings.TrimSpace(removePortFromRawAddr(r.RemoteAddr)),
		"command": strings.TrimSpace(r.Method),
		"real_path": strings.TrimSpace(r.URL.EscapedPath()),
		"query": strings.TrimSpace(r.URL.Query().Encode()),
		"version": strings.TrimSpace(r.Proto),
		"request_uri": strings.TrimSpace(r.RequestURI),
	}
}

func removePortFromRawAddr(remoteAddr string) string {
	if index := strings.LastIndexByte(remoteAddr, ':'); index != -1 {
		remoteAddr = remoteAddr[:index]
	}

	return remoteAddr
}

func IsFromCurl(r *http.Request) bool {
	userAgent := r.Header.Get("User-Agent")
	return strings.HasPrefix(userAgent, "curl/")
}
