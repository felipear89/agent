package server

import (
	"log/slog"

	"github.com/felipear89/agent/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) setupMiddleware() {
	s.router.Use(gin.Recovery())
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Timeout(middleware.TimeoutConfig{
		Timeout:      s.config.Timeout,
		ErrorMessage: "Request processing timed out",
		OnTimeout: func(c *gin.Context) {
			slog.ErrorContext(c.Request.Context(), "Request timeout handler triggered",
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
			)
		},
	}))
}
