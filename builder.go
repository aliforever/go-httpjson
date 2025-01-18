package httpjson

import (
	"encoding/json"
	"net/http"
)

type Builder struct {
	message string
	data    interface{}
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetMessage(message string) *Builder {
	b.message = message
	return b
}

func (b *Builder) SetData(data interface{}) *Builder {
	b.data = data

	return b
}

func (b *Builder) JSON(statusCode int) ([]byte, error) {
	return json.Marshal(newResponse(statusCode, b.message, b.data))
}

func (b *Builder) Ok(writer http.ResponseWriter) error {
	return Ok(writer, b.message, b.data)
}

func (b *Builder) Accepted(writer http.ResponseWriter) error {
	return Accepted(writer, b.message, b.data)
}

func (b *Builder) BadRequest(writer http.ResponseWriter) error {
	return BadRequest(writer, b.message, b.data)
}

func (b *Builder) Unauthorized(writer http.ResponseWriter) error {
	return Unauthorized(writer, b.message, b.data)
}

func (b *Builder) Forbidden(writer http.ResponseWriter) error {
	return Forbidden(writer, b.message, b.data)
}

func (b *Builder) NotFound(writer http.ResponseWriter) error {
	return NotFound(writer, b.message, b.data)
}

func (b *Builder) MethodNotAllowed(writer http.ResponseWriter) error {
	return MethodNotAllowed(writer, b.message, b.data)
}

func (b *Builder) InternalServerError(writer http.ResponseWriter) error {
	return InternalServerError(writer, b.message, b.data)
}
