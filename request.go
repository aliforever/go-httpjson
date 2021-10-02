package httpjson

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	action string
	data   json.RawMessage
}

func ParseRequest(wr http.ResponseWriter, r *http.Request) (req *Request, err error) {
	type Req struct {
		Action string          `json:"action"`
		Data   json.RawMessage `json:"data"`
	}

	var publicReq *Req
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&publicReq)
	if err != nil {
		BadRequest(wr, "invalid_json")
		return
	}
	req = &Request{
		action: publicReq.Action,
		data:   publicReq.Data,
	}
	return
}

func (r *Request) Action() string {
	return r.action
}

func (r *Request) UnmarshalData(dst interface{}) (err error) {
	err = json.Unmarshal(r.data, &dst)
	return
}
