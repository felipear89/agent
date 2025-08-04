package user

import "github.com/gin-gonic/gin"

func Register(api *gin.RouterGroup) {
	repository := NewInMemoryRepository()
	service := NewService(repository)
	handler := NewHandler(service)
	handler.RegisterRoutes(api)
}
