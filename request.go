package httpjson

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

func ParseRequest(wr http.ResponseWriter, r *http.Request) (req *Request, err error) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		BadRequest(wr, "invalid_json")
	}
	return
}
