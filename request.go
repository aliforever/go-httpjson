package httpjson

import (
	"encoding/json"
	"net/http"
)

type request struct {
	payload

	Data json.RawMessage `json:"data,omitempty"`
}

func ParsePayloadData[T any](r *http.Request) (t *T, err error) {
	defer r.Body.Close()

	var p *request

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(p.Data, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func UnmarshalRequest(r *http.Request, destination interface{}) (err error) {
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&destination)

	return
}
