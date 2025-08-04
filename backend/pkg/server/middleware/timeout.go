package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	apperrors "github.com/felipear89/agent/pkg/server/errors"
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
		OnTimeout: func(c *gin.Context) {
			slog.ErrorContext(c.Request.Context(), "Request timeout",
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
			)
		},
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
					c.Error(apperrors.New(
						apperrors.ErrCodeTimeout,
						config.ErrorMessage,
						http.StatusRequestTimeout,
					))
					c.Abort()
					return
				}
				c.Abort()
			}
		}
	}
}
