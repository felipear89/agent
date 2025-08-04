package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"

	apperrors "github.com/felipear89/agent/pkg/server/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *apperrors.AppError

			switch {
			case errors.As(err, &appErr):
				c.JSON(appErr.Status, gin.H{
					"error": gin.H{
						"code":    appErr.Code,
						"message": appErr.Message,
					},
				})
			case err != nil:
				// Handle unexpected errors
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{
						"code":    apperrors.ErrCodeInternal,
						"message": "An unexpected error occurred",
					},
				})
			}
		}
	}
}
