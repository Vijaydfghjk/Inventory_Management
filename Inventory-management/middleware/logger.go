package middleware

import (
	token_stuff "inventory_management/Token_stuff"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := token_stuff.ValidateJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)

		// Role check

		if len(allowedRoles) == 0 {
			c.Next()
			return
		}

		for _, role := range allowedRoles {
			if claims.Role == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}
