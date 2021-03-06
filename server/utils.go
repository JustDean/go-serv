package server

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseJSON(request *http.Request) (map[string]string, error) {
	var res map[string]string
	b, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return nil, nil //TODO handle case when body is empty
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ReturnJSON(w http.ResponseWriter, content map[string]string) error {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		return err
	}

	return nil
}
