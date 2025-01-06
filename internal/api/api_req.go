package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseJSON(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer func() {
		_ = r.Body.Close()
	}()
	return json.Unmarshal(body, v)
}
