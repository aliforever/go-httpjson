package httpjson

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type request struct {
	Payload

	Data json.RawMessage `json:"data,omitempty"`
}

func ParsePayloadData[T any](wr http.ResponseWriter, r *http.Request) (t *T, err error) {
	defer r.Body.Close()

	var p *request

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeResponseErr := BadRequest(wr, "invalid_json", nil)
		if writeResponseErr != nil {
			return nil, fmt.Errorf("parse_error: %s - write_response_err: %s", err, writeResponseErr)
		}
		return nil, err
	}

	err = json.Unmarshal(p.Data, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func UnmarshalRequest(wr http.ResponseWriter, r *http.Request, destination interface{}) (err error) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&destination)
	if err != nil {
		err := BadRequest(wr, "invalid_json", nil)
		if err != nil {
			return err
		}
		return
	}
	return
}
