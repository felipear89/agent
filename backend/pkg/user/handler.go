package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
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
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    users,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to fetch user",
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    user,
	})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "User created successfully",
		Data:    createdUser,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user.ID = id
	updatedUser, err := h.service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	if updatedUser == nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "User updated successfully",
		Data:    updatedUser,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "User deleted successfully",
	})
}
