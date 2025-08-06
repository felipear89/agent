package middleware

import (
	"github.com/felipear89/agent/pkg/server/apperror"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func Recovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(gin.DefaultErrorWriter, func(c *gin.Context, err any) {

		slog.ErrorContext(c.Request.Context(), "Recovered from panic", "error", err, "path", c.Request.URL.Path)

		switch e := err.(type) {
		case string:
			apperror.InternalErrorResponse(c, e)
		case error:
			apperror.InternalErrorCustomResponse(c, e, "Internal server error")
		default:
			apperror.InternalErrorResponse(c, "Unknown panic")
		}
	})
}
