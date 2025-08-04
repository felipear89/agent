package middleware

import (
	"log/slog"
	"net/http"

	apperrors "github.com/felipear89/agent/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(gin.DefaultErrorWriter, func(c *gin.Context, err any) {
		slog.ErrorContext(c.Request.Context(), "Recovered from panic",
			"error", err,
			"path", c.Request.URL.Path,
		)
		c.Error(apperrors.New(
			apperrors.ErrCodeInternal,
			"Internal server error",
			http.StatusInternalServerError,
		))
		c.Abort()
		return
	})
}
