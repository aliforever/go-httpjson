package httpjson

import (
	"encoding/json"
	"net/http"
)

func ParseRequest(wr http.ResponseWriter, r *http.Request, destination interface{}) (err error) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&destination)
	if err != nil {
		BadRequest(wr, "invalid_json")
		return
	}
	return
}
