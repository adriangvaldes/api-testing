package middlewares

import (
	"net/http"
	"strings"

	"github.com/adriangvaldes/api-testing/helper"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		// Check if the token is missing
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the Authorization header has the correct format (Bearer <token>)
		token := strings.Split(authorization, " ")
		if len(token) != 2 || token[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Verify token
		userId, err := helper.VerifyToken(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}
