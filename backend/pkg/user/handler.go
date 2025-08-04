package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/felipear89/agent/pkg/server/errors"
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
		group.PUT("/:id", h.UpdateUser)
		group.DELETE("/:id", h.DeleteUser)
	}
}

func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.Error(errors.Wrap(err, errors.ErrCodeInternal, "failed to fetch users", http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errors.New(
			errors.ErrCodeInvalidInput,
			"Invalid user ID",
			http.StatusBadRequest,
		))
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		c.Error(errors.Wrap(
			err,
			errors.ErrCodeNotFound,
			fmt.Sprintf("User with ID %d not found", id),
			http.StatusNotFound,
		))
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
		c.Error(errors.New(
			errors.ErrCodeInvalidInput,
			"Invalid request body",
			http.StatusBadRequest,
		))
		return
	}

	// Validate request
	if req.Name == "" || req.Email == "" {
		c.Error(errors.New(
			errors.ErrCodeInvalidInput,
			"Name and email are required",
			http.StatusBadRequest,
		))
		return
	}

	user, err := h.service.CreateUser(User{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		c.Error(errors.Wrap(
			err,
			errors.ErrCodeInternal,
			"Failed to create user",
			http.StatusInternalServerError,
		))
		return
	}

	c.JSON(http.StatusCreated, user)
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errors.New(
			errors.ErrCodeInvalidInput,
			"Invalid user ID",
			http.StatusBadRequest,
		))
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.New(
			errors.ErrCodeInvalidInput,
			"Invalid request body",
			http.StatusBadRequest,
		))
		return
	}

	// Validate request
	if req.Name == "" && req.Email == "" {
		c.Error(errors.New(
			errors.ErrCodeInvalidInput,
			"At least one field (name or email) is required",
			http.StatusBadRequest,
		))
		return
	}

	updatedUser, err := h.service.UpdateUser(User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		c.Error(errors.Wrap(
			err,
			errors.ErrCodeInternal,
			"Failed to update user",
			http.StatusInternalServerError,
		))
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errors.New(
			errors.ErrCodeInvalidInput,
			"Invalid user ID",
			http.StatusBadRequest,
		))
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		c.Error(errors.Wrap(
			err,
			errors.ErrCodeInternal,
			"Failed to delete user",
			http.StatusInternalServerError,
		))
		return
	}

	c.Status(http.StatusNoContent)
}
