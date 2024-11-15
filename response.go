package httpjson

import (
	"encoding/json"
	"net/http"
)

type response struct {
	payload

	Data interface{} `json:"data,omitempty"`
}

func newResponse(statusCode int, message string, data interface{}) response {
	return response{
		payload: payload{
			StatusCode: statusCode,
			StatusName: statusMessage(statusCode),
			Message:    message,
		},
		Data: data,
	}
}

func writeHeadersAndData(writer http.ResponseWriter, statusCode int, message string, data interface{}) (err error) {
	var jsonBytes []byte

	jsonBytes, err = json.Marshal(newResponse(statusCode, message, data))
	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	_, err = writer.Write(jsonBytes)

	return err
}

func writeHeadersAndDataCustom(writer http.ResponseWriter, statusCode int, data interface{}) (err error) {
	var jsonBytes []byte

	jsonBytes, err = json.Marshal(data)
	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	_, err = writer.Write(jsonBytes)

	return err
}
