package httpjson

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	StatusCode int         `json:"status_code"`
	StatusName string      `json:"status_name"`
	Data       interface{} `json:"data,omitempty"`
}

func writeHeadersAndData(writer http.ResponseWriter, statusCode int, isRaw bool, data interface{}) (err error) {
	if !isRaw {
		var jsonBytes []byte
		jsonBytes, err = json.Marshal(response{
			StatusCode: statusCode,
			StatusName: statusMessage(statusCode),
			Data:       data,
		})
		if returnOnJsonMarshalError(writer, err) {
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(statusCode)
		writer.Write(jsonBytes)
	} else {
		writer.Header().Set("Content-Type", "application/text")
		writer.WriteHeader(statusCode)

		var r []byte

		switch data.(type) {
		case string:
			r = []byte(data.(string))
			break
		case uint, uint8, uint16, uint32, uint64, int, int8, int32, int64:
			r = []byte(fmt.Sprintf("%d", data.(int)))
			break
		case float32, float64:
			r = []byte(fmt.Sprintf("%f", data))
			break
		case bool:
			r = []byte(fmt.Sprintf("%t", data.(bool)))
			break
		default:
			r, _ = json.Marshal(data)
			if r == nil {
				// TODO: It might be better to return Internal Server Error and report unknown types
				r = []byte("EMPTY_RESPONSE")
			}
		}
		writer.Write(r)
	}
	return
}

func returnOnJsonMarshalError(writer http.ResponseWriter, err error) (shouldReturn bool) {
	if err != nil {
		shouldReturn = true
		message := "Internal Server Error"
		writeHeadersAndData(writer, http.StatusInternalServerError, true, &message)
	}
	return
}
