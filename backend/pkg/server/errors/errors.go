package errors

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorCode string

const (
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

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *AppError

			switch {
			case errors.As(err, &appErr):
				c.JSON(appErr.Status, gin.H{
					"error": gin.H{
						"code":    appErr.Code,
						"message": appErr.Message,
					},
				})
			case err != nil:
				// Handle unexpected errors
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{
						"code":    ErrCodeInternal,
						"message": "An unexpected error occurred",
					},
				})
			}
		}
	}
}

func NewInternalError(err error) *AppError {
	return Wrap(err, ErrCodeInternal, "Internal server error", http.StatusInternalServerError)
}

func NewNotFoundError(resource string) *AppError {
	return New(ErrCodeNotFound, fmt.Sprintf("%s not found", resource), http.StatusNotFound)
}

func NewValidationError(message string) *AppError {
	return New(ErrCodeInvalidInput, message, http.StatusBadRequest)
}

func NewUnauthorizedError(message string) *AppError {
	if message == "" {
		message = "You are not authorized to access this resource"
	}
	return New(ErrCodeUnauthorized, message, http.StatusUnauthorized)
}
