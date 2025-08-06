package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/felipear89/agent/pkg/server/apperror"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterAPIRoutes() *gin.RouterGroup {
	api := s.router.Group(s.config.BasePath)
	health(api)
	slow(api)
	api.GET("/test-panic", func(c *gin.Context) {
		var nilMap map[string]string
		// This will cause a panic
		nilMap["key"] = "value"
	})

	api.GET("/bad-request", func(c *gin.Context) {
		apperror.BadRequestResponse(c, errors.New("Bad request example"))
		return
	})
	return api
}

// Delay returns nil after the specified duration or error if interrupted.
func Delay(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	select {
	case <-ctx.Done():
		t.Stop()
		return ctx.Err()
	case <-t.C:
	}
	return nil
}

func slow(api *gin.RouterGroup) gin.IRoutes {

	return api.GET("/slow", func(c *gin.Context) {
		err := Delay(c.Request.Context(), 8*time.Second)
		if err != nil {
			apperror.InternalErrorCustomResponse(c, err, "slow endpoint")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "This will never be reached"})
	})
}

func health(api *gin.RouterGroup) gin.IRoutes {
	return api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().UTC().Format(time.RFC3339),
		})
	})
}
