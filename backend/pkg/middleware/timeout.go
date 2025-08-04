package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TimeoutConfig struct {
	Timeout      time.Duration
	ErrorMessage string
	OnTimeout    func(*gin.Context)
}

func DefaultTimeoutConfig() TimeoutConfig {
	return TimeoutConfig{
		Timeout:      10 * time.Second,
		ErrorMessage: "Request processing timed out",
		OnTimeout:    nil,
	}
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

		slog.Info("Starting timeout middleware")

		ctx, cancel := context.WithTimeout(c.Request.Context(), config.Timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan struct{})

		go func() {
			defer close(done)
			c.Next()
		}()

		select {
		case <-done:
			return

		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				slog.ErrorContext(c.Request.Context(), "Request processing timed out",
					"method", c.Request.Method,
					"path", c.Request.URL.Path,
					"timeout", config.Timeout.String(),
				)

				if config.OnTimeout != nil {
					config.OnTimeout(c)
				}

				if !c.Writer.Written() {
					c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
						"error": config.ErrorMessage,
					})
				}
				c.Abort()
			}
		}
	}
}
