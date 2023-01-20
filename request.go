package httpjson

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	payload

	Data json.RawMessage `json:"data,omitempty"`
}

func UnmarshalRequest[T any](r *http.Request) (t *T, err error) {
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&t)

	return
}

// UnmarshalRequestPayload this is to unmarshal the request using the library's payload structure
func UnmarshalRequestPayload[T any](r *http.Request) (t *T, err error) {
	defer r.Body.Close()

	var request *Request
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(request.Data, &t)
	if err != nil {
		return nil, err
	}

	return
}

func UnmarshalFromBytes[T any](data []byte) (t *T, err error) {
	err = json.Unmarshal(data, &t)

	return
}

// UnmarshalPayloadFromBytes this is to unmarshal the request using the library's payload structure
func UnmarshalPayloadFromBytes[T any](data []byte) (t *T, err error) {
	var request *Request
	err = json.Unmarshal(data, &request)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(request.Data, &t)
	if err != nil {
		return nil, err
	}

	return
}

func Post[Response any](address string, data interface{}) (*Response, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(address, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r *Response

	err = json.NewDecoder(resp.Body).Decode(&r)

	return r, err
}

func PostFromReader[Response any](address string, data io.Reader) (*Response, error) {
	resp, err := http.Post(address, "application/json", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r *Response

	err = json.NewDecoder(resp.Body).Decode(&r)

	return r, err
}
