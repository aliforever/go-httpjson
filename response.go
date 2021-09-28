package httpjson

import (
	"encoding/json"
	"net/http"
)

type response struct {
	StatusCode int         `json:"status_code"`
	StatusName string      `json:"status_name"`
	Data       interface{} `json:"data,omitempty"`
}

func writeHeadersAndData(writer http.ResponseWriter, statusCode int, data interface{}) (err error) {
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(response{
		StatusCode: statusCode,
		StatusName: statusMessage(statusCode),
		Data:       data,
	})
	if err != nil {
		writer.Header().Set("Content-Type", "application/text")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(jsonBytes)
	return
}
