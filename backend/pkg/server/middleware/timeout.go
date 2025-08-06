package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type TimeoutConfig struct {
	Timeout      time.Duration
	ErrorMessage string
	OnTimeout    func(*gin.Context)
}

func DefaultTimeout(duration time.Duration) gin.HandlerFunc {
	return Timeout(TimeoutConfig{
		Timeout:      duration,
		ErrorMessage: "Request processing timed out",
	})
}

func Timeout(config TimeoutConfig) gin.HandlerFunc {
	if config.Timeout <= 0 {
		config.Timeout = 10 * time.Second
	}

	if config.ErrorMessage == "" {
		config.ErrorMessage = "Request processing timed out"
	}

	return func(c *gin.Context) {
		if c.Request.URL.Path == "/stream" {
			c.Next()
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), config.Timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
