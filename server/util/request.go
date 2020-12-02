package util

import (
	"bytes"
	"net/http"
)

func GetRequestAsString(r *http.Request) string {
	buf := new(bytes.Buffer)
	if err := r.Write(buf); err != nil {
		return err.Error()
	}
	return string(buf.Bytes())
}
