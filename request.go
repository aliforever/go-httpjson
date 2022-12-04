package httpjson

import (
	"encoding/json"
	"net/http"
)

func UnmarshalRequest[T any](r *http.Request) (t *T, err error) {
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&t)

	return
}
