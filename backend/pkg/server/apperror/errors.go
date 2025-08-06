package apperror

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
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

func (e ErrorCode) Error() string {
	return string(e)
}

type ErrorResponse struct {
	Error AppErrorDetail `json:"error"`
}

type AppErrorDetail struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func New(code ErrorCode, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: AppErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}

func InternalErrorResponse(c *gin.Context, msg string) {
	abort(c, http.StatusInternalServerError, New(ErrCodeInternal, msg))
}

func InternalErrorCustomResponse(c *gin.Context, err error, msg string) {
	c.Error(err)
	abort(c, http.StatusInternalServerError, New(ErrCodeInternal, msg))
}

func BadRequestResponse(c *gin.Context, err error) {
	c.Error(err)
	abort(c, http.StatusBadRequest, New(ErrCodeInvalidInput, "Invalid request body"))
}

func BadRequestCustomResponse(c *gin.Context, err error, msg string) {
	c.Error(err)
	abort(c, http.StatusBadRequest, New(ErrCodeInvalidInput, msg))
}

func NotFoundResponse(c *gin.Context, err error, msg string) {
	c.Error(err)
	abort(c, http.StatusNotFound, New(ErrCodeNotFound, msg))
}

func UnauthorizedResponse(c *gin.Context, msg string) {
	abort(c, http.StatusUnauthorized, New(ErrCodeUnauthorized, msg))
}

func abort(c *gin.Context, status int, response *ErrorResponse) {
	if errors.Is(c.Errors.Last(), context.DeadlineExceeded) {
		c.AbortWithStatusJSON(http.StatusRequestTimeout, New(ErrCodeTimeout, "Request processing timed out"))
		return
	}
	c.AbortWithStatusJSON(status, response)
}
