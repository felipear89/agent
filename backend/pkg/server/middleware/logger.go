package middleware

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		logger := slog.With(
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"query", c.Request.URL.RawQuery,
			"ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"status", status,
			"latency", duration.String(),
			"time", time.Now().UTC().Format(time.RFC3339),
		)

		msg := "Request processed"
		switch {
		case status >= 500:
			logger.ErrorContext(c.Request.Context(), msg, "errors", c.Errors.ByType(gin.ErrorTypePrivate).String())
		case status >= 400:
			logger.WarnContext(c.Request.Context(), msg, "errors", c.Errors.ByType(gin.ErrorTypePrivate).String())
		default:
			logger.InfoContext(c.Request.Context(), msg)
		}
	}
}
