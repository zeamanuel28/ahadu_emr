package middleware

import (
	"net/http"
	"strings"

	"saas-api/shared/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization header required", "")
			c.Abort()
			return
		}

		token := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else if strings.Contains(authHeader, " ") {
			// If it contains a space but doesn't start with Bearer, it's likely invalid
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization format. Use 'Bearer <token>' or just the token.", "")
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token", err.Error())
			c.Abort()
			return
		}

		// Set user info to context
		c.Set("userId", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("roles", claims.Roles)

		c.Next()
	}
}
