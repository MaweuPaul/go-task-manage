package middlewares

import (
	"os"
	"taskmanager/utils"

	"github.com/gin-gonic/gin"
)

// Middleware to protect routes - checks for valid JWT in cookies
func AuthMiddleware(c *gin.Context) {

	accessToken, err := c.Cookie("accessToken")

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return

	}

	claims, err := utils.ValidateToken(accessToken)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
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
	//check if app i slive in production or development
	c.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_ORIGINS"))
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Next()
}

//
