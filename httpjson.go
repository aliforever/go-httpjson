package httpjson

import (
	"net/http"
)

func Ok(writer http.ResponseWriter, message string, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusOK, message, data)
	return
}

func Accepted(writer http.ResponseWriter, message string, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusAccepted, message, data)
	return
}

func BadRequest(writer http.ResponseWriter, message string, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusBadRequest, message, data)
	return
}

func Unauthorized(writer http.ResponseWriter, message string, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusUnauthorized, message, data)
	return
}

func Forbidden(writer http.ResponseWriter, message string, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusForbidden, message, data)
	return
}

func NotFound(writer http.ResponseWriter, message string, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusNotFound, message, data)
	return
}

func MethodNotAllowed(writer http.ResponseWriter, message string, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusMethodNotAllowed, message, data)
	return
}

func InternalServerError(writer http.ResponseWriter, message string, data interface{}) (err error) {
	err = writeHeadersAndData(writer, http.StatusInternalServerError, message, data)
	return
}
