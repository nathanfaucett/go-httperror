package httperror

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Code    int
	Message string
}

func New(code int, message string) *HttpError {
	this := new(HttpError)

	this.Code = code
	this.Message = message

	return this
}

func NewCode(code int) *HttpError {
	this := new(HttpError)

	this.Code = code

	return this
}

func NewMessage(message string) *HttpError {
	this := new(HttpError)

	this.Code = 500
	this.Message = message

	return this
}

func (this *HttpError) Error() string {
	if this.Message != "" {
		return fmt.Sprintf("%d %s - %s", this.Code, http.StatusText(this.Code), this.Message)
	}

	return fmt.Sprintf("%d %s", this.Code, http.StatusText(this.Code))
}
