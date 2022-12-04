package tests

import (
	"bytes"
	"encoding/json"
	"github.com/aliforever/go-httpjson"
	"net/http"
	"testing"
)

func TestParsePayloadData(t *testing.T) {
	exit := make(chan bool)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			exit <- true
		}()
		type data struct {
			ID int64 `json:"id"`
		}

		d, err := httpjson.UnmarshalRequest[data](request)
		if err != nil {
			_ = httpjson.BadRequest(writer, err.Error(), nil)
			t.Fatal("invalid_request", err)
		}

		if d.ID != 1 {
			t.Fatal("invalid value for id")
		}
	})
	go http.ListenAndServe(":80", nil)

	j, _ := json.Marshal(map[string]interface{}{
		"id": 1,
	})

	go func() {
		_, err := http.Post("http://127.0.0.1:80", "application/json", bytes.NewReader(j))
		if err != nil {
			t.Error(err)
			return
		}
	}()

	<-exit
}
