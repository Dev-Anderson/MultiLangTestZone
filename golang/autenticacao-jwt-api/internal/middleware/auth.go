package middleware

import (
	"aut-jwt/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
