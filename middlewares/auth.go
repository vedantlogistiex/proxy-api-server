package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CheckHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := os.Getenv("WPROXY_KEY")
		headerValue := c.GetHeader("X-WPROXY-KEY")
		if headerValue != key {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}
