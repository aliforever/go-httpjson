package httpjson

import (
	"net/http"
)

func Custom(writer http.ResponseWriter, statusCode int, data interface{}) (err error) {
	err = writeHeadersAndDataCustom(writer, statusCode, data)
	return
}

func OkCustom(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndDataCustom(writer, http.StatusOK, data)
	return
}

func AcceptedCustom(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndDataCustom(writer, http.StatusAccepted, data)
	return
}

func BadRequestCustom(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndDataCustom(writer, http.StatusBadRequest, data)
	return
}

func UnauthorizedCustom(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndDataCustom(writer, http.StatusUnauthorized, data)
	return
}

func ForbiddenCustom(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndDataCustom(writer, http.StatusForbidden, data)
	return
}

func NotFoundCustom(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndDataCustom(writer, http.StatusNotFound, data)
	return
}

func MethodNotAllowedCustom(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndDataCustom(writer, http.StatusMethodNotAllowed, data)
	return
}

func InternalServerErrorCustom(writer http.ResponseWriter, data interface{}) (err error) {
	err = writeHeadersAndDataCustom(writer, http.StatusInternalServerError, data)
	return
}
