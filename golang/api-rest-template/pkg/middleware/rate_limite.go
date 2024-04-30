package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter(r rate.Limit, b int) gin.HandlerFunc {
	limiter := rate.NewLimiter(r, b)
	return func(c *gin.Context) {
		if !limiter.AllowN(time.Now(), 1) {
			c.String(http.StatusTooManyRequests, "Limite de acesso excedido")
			c.Abort()
			return
		}
		c.Next()
	}
}
