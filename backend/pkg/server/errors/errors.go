package errors

import (
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	ErrCodeTimeout      ErrorCode = "TIMEOUT"
	ErrCodeInternal     ErrorCode = "INTERNAL_ERROR"
	ErrCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrCodeInvalidInput ErrorCode = "INVALID_INPUT"
	ErrCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden    ErrorCode = "FORBIDDEN"
)

type AppError struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	InnerError error     `json:"-"`
	Status     int       `json:"-"`
}

func (e *AppError) Error() string {
	if e.InnerError != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.InnerError)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.InnerError
}

func New(code ErrorCode, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

func Wrap(err error, code ErrorCode, message string, status int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		InnerError: err,
		Status:     status,
	}
}

func NewInternalError(err error) *AppError {
	return Wrap(err, ErrCodeInternal, "Internal server error", http.StatusInternalServerError)
}
