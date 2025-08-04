package auth

import (
	"github.com/felipear89/agent/pkg/user"
	"net/http"

	"github.com/felipear89/agent/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type AuthHandler struct {
	authService *Service
	userService *user.Service
}

func NewHandler(api *gin.RouterGroup, authService *Service, userService *user.Service) *AuthHandler {
	h := &AuthHandler{
		authService: authService,
		userService: userService,
	}
	h.RegisterRoutes(api)
	return h
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.New(
			errors.ErrCodeInvalidInput,
			"Invalid request body",
			http.StatusBadRequest,
		))
		return
	}

	// Authenticate user
	user, err := h.userService.Authenticate(req.Email, req.Password)
	if err != nil {
		c.Error(errors.New(
			errors.ErrCodeUnauthorized,
			"Invalid email or password",
			http.StatusUnauthorized,
		))
		return
	}

	// Generate JWT token
	token, err := h.authService.GenerateToken(user.ID, req.Email)
	if err != nil {
		c.Error(errors.Wrap(
			err,
			errors.ErrCodeInternal,
			"Failed to generate token",
			http.StatusInternalServerError,
		))
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int64(h.authService.cfg.TokenExpiryDuration().Seconds()),
	})
}

// RegisterRoutes registers auth routes
func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", h.Login)
}
