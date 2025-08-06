package auth

import (
	"strings"

	"github.com/felipear89/agent/pkg/server/apperror"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	BearerTokenType     = "Bearer"
	UserId              = "user_id"
	userEmail           = "user_email"
)

// AuthMiddleware verifies the JWT token in the Authorization header
func (a *Service) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			apperror.UnauthorizedResponse(c, "Authorization header is required")
			return
		}

		// Extract the token from the header
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != BearerTokenType {
			apperror.UnauthorizedResponse(c, "Authorization header is invalid")
			return
		}

		tokenString := tokenParts[1]
		claims, err := a.ValidateToken(tokenString)
		if err != nil {
			apperror.UnauthorizedResponse(c, "Invalid or expired token")
			return
		}

		// Add user info to context
		c.Set(UserId, claims.UserID)
		c.Set(userEmail, claims.Email)

		c.Next()
	}
}

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
