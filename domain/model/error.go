package model

import (
	"fmt"
	"net/http"
)

var (
	ErrInvalidSignature = NewError(http.StatusBadRequest, "Invalid signature")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Message)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
