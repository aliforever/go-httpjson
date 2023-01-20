package httpjson

import (
	"encoding/json"
	"net/http"
)

type Builder struct {
	writer  http.ResponseWriter
	message string
	data    interface{}
}

func New(writer http.ResponseWriter) Builder {
	return Builder{writer: writer}
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

func (b *Builder) Ok() error {
	return Ok(b.writer, b.message, b.data)
}

func (b *Builder) Accepted() error {
	return Accepted(b.writer, b.message, b.data)
}

func (b *Builder) BadRequest() error {
	return BadRequest(b.writer, b.message, b.data)
}

func (b *Builder) Unauthorized() error {
	return Unauthorized(b.writer, b.message, b.data)
}

func (b *Builder) Forbidden() error {
	return Forbidden(b.writer, b.message, b.data)
}

func (b *Builder) NotFound() error {
	return NotFound(b.writer, b.message, b.data)
}

func (b *Builder) MethodNotAllowed() error {
	return MethodNotAllowed(b.writer, b.message, b.data)
}

func (b *Builder) InternalServerError() error {
	return InternalServerError(b.writer, b.message, b.data)
}
