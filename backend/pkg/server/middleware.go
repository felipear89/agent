package server

import (
	"log/slog"
	"net/http"

	"github.com/felipear89/agent/pkg/server/errors"
	"github.com/felipear89/agent/pkg/server/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) setupMiddleware() {
	// Recovery middleware to handle panics
	s.router.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter, func(c *gin.Context, err any) {
		slog.ErrorContext(c.Request.Context(), "Recovered from panic",
			"error", err,
			"path", c.Request.URL.Path,
		)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"code":    errors.ErrCodeInternal,
				"message": "Internal server error",
			},
		})
	}))

	// Request logging
	s.router.Use(middleware.Logger())

	// Timeout handling
	s.router.Use(middleware.Timeout(middleware.TimeoutConfig{
		Timeout:      s.config.Timeout,
		ErrorMessage: "Request processing timed out",
		OnTimeout: func(c *gin.Context) {
			slog.ErrorContext(c.Request.Context(), "Request timeout",
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
			)
		},
	}))

	// Global error handler (must be registered after all other middleware)
	s.router.Use(errors.ErrorHandler())
}
