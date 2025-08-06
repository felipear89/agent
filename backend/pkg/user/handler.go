package user

import (
	"net/http"
	"strconv"

	"github.com/felipear89/agent/pkg/server/apperror"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(api *gin.RouterGroup, service *Service) *Handler {
	h := &Handler{service: service}
	h.RegisterRoutes(api)
	return h
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/users")
	{
		group.GET("", h.ListUsers)
		group.POST("", h.CreateUser)
		group.GET("/:id", h.GetUser)
	}
}

func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		apperror.InternalErrorCustomResponse(c, err, "Failed to fetch users")
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apperror.BadRequestCustomResponse(c, err, "Invalid user ID format")
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		apperror.NotFoundResponse(c, err, "User not found")
		return
	}

	c.JSON(http.StatusOK, user)
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.BadRequestCustomResponse(c, err, err.Error())
		return
	}

	// Validate request
	if req.Name == "" || req.Email == "" {
		apperror.BadRequestCustomResponse(c, nil, "Name and email are required")
		return
	}

	user, err := h.service.CreateUser(User{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		apperror.InternalErrorCustomResponse(c, err, "Failed to create user")
		return
	}

	c.JSON(http.StatusCreated, user)
}
