package auth

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/felipear89/agent/pkg/server/errors"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	BearerTokenType     = "Bearer"
)

// AuthMiddleware verifies the JWT token in the Authorization header
func (a *Service) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			c.Error(errors.New(
				errors.ErrCodeUnauthorized,
				"Authorization header is required",
				http.StatusUnauthorized,
			))
			c.Abort()
			return
		}

		// Extract the token from the header
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != BearerTokenType {
			c.Error(errors.New(
				errors.ErrCodeUnauthorized,
				"Invalid authorization header format",
				http.StatusUnauthorized,
			))
			c.Abort()
			return
		}

		tokenString := tokenParts[1]
		claims, err := a.ValidateToken(tokenString)
		if err != nil {
			slog.Error("Invalid token", "error", err)
			c.Error(errors.New(
				errors.ErrCodeUnauthorized,
				"Invalid or expired token",
				http.StatusUnauthorized,
			))
			c.Abort()
			return
		}

		// Add user info to context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}

// GetUserIDFromContext retrieves the user ID from the context
func GetUserIDFromContext(c *gin.Context) (int, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	id, ok := userID.(int)
	if !ok {
		return 0, false
	}

	return id, true
}

// GetUserEmailFromContext retrieves the user email from the context
func GetUserEmailFromContext(c *gin.Context) (string, bool) {
	email, exists := c.Get("user_email")
	if !exists {
		return "", false
	}

	emailStr, ok := email.(string)
	if !ok {
		return "", false
	}

	return emailStr, true
}
