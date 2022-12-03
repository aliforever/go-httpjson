package httpjson

import (
	"encoding/json"
	"net/http"
)

type response struct {
	payload

	Data interface{} `json:"data,omitempty"`
}

func writeHeadersAndData(writer http.ResponseWriter, statusCode int, message string, data interface{}) (err error) {
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(response{
		payload: payload{
			StatusCode: statusCode,
			StatusName: statusMessage(statusCode),
			Message:    message,
		},
		Data: data,
	})
	if err != nil {
		writer.Header().Set("Content-Type", "application/text")
		writer.WriteHeader(http.StatusInternalServerError)
		_, err := writer.Write([]byte("Internal Server Error"))
		if err != nil {
			return err
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	_, err = writer.Write(jsonBytes)
	if err != nil {
		return err
	}
	return
}
