package server

import (
	"net/http"
	"time"

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
	return api
}

func slow(api *gin.RouterGroup) gin.IRoutes {
	return api.GET("/slow", func(c *gin.Context) {
		time.Sleep(20 * time.Second)
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
