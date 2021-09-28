package httpjson

import (
	"net/http"
)

func Ok(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusOK, data)
	return
}

func Accepted(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusAccepted, data)
	return
}

func BadRequest(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusBadRequest, data)
	return
}

func Unauthorized(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusUnauthorized, data)
	return
}

func Forbidden(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusForbidden, data)
	return
}

func NotFound(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusNotFound, data)
	return
}

func MethodNotAllowed(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusMethodNotAllowed, data)
	return
}
