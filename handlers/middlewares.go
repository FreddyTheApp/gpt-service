package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h GinHandler) SecretMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestSecret := c.GetHeader("X-Secret")

		if requestSecret == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No secret in the header"})
			c.Abort()
			return
		}

		envSecret := os.Getenv("API_SECRET")

		if requestSecret != envSecret {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid secret"})
			c.Abort()
			return
		}

		c.Next()
	}
}
