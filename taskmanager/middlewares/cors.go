package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	if allowedOrigins != "" && strings.Contains(allowedOrigins, origin) {
		c.Header("Access-Control-Allow-Origin", origin)
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}
