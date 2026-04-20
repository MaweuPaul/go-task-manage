package middlewares

import (
	"net/http"
	"os"
	"strings"
	"taskmanager/utils"

	"github.com/gin-gonic/gin"
)

// Middleware to protect routes - checks for valid JWT in cookies
func AuthMiddleware(c *gin.Context) {
	accessToken, err := c.Cookie("accessToken")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	claims, err := utils.ValidateToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid or expired token",
		})
		return
	}

	c.Set("userId", claims["userId"])
	c.Next()
}

// rate limiting
func RateLimitMiddleware(c *gin.Context) {

}

// cors handling
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

//
