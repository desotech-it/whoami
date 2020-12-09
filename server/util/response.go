package util

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func WriteJSONResponse(w http.ResponseWriter, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	header := w.Header()
	header.Set("Content-Type", "application/json; charset=utf-8")
	header.Set("Content-Length", strconv.Itoa(len(bytes)))
	_, errWrite := w.Write(bytes)
	return errWrite
}
