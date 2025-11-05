package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kshzz24/gosocial/internal/utils"
)

func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth_header := c.GetHeader("Token")
		if auth_header == "" {
			c.Set("user_id", nil)
			c.Set("is_authenticated", false)
			c.Next()
			return
		}
		parts := strings.Split(auth_header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {

			c.Set("user_id", nil)
			c.Set("is_authenticated", false)
			c.Next()
			return
		}
		token := parts[1]

		// Validate token
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			c.Set("user_id", nil)
			c.Set("is_authenticated", false)
			c.Next()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)
		c.Set("is_authenticated", true)
		c.Next()
	}
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Token")

		// No header = unauthorized
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is  required"})
			c.Abort()
			return
		}

		token := authHeader

		// Validate token
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Valid token = set user info and continue
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("is_authenticated", true)

		c.Next()
	}
}
