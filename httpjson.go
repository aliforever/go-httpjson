package httpjson

import (
	"net/http"
)

func Ok(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusOK, false, data)
	return
}

func Accepted(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusAccepted, false, &data)
	return
}

func BadRequest(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusBadRequest, false, &data)
	return
}

func Unauthorized(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusUnauthorized, false, &data)
	return
}

func Forbidden(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusForbidden, false, &data)
	return
}

func NotFound(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusNotFound, false, &data)
	return
}

func MethodNotAllowed(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusMethodNotAllowed, false, &data)
	return
}
