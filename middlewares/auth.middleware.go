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

// Middleware to refresh access token if expired - checks refresh token in cookies
func RefreshTokenMiddleware(c *gin.Context) {

	domain := os.Getenv("DOMAIN")

	refreshToken, err := c.Cookie("refreshToken")

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	claims, err := utils.ValidateToken(refreshToken)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userId := claims["userId"].(string)

	newAccessToken, err := utils.GenerateJWT(userId)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to generate access token"})
		return
	}
	c.SetCookie("accessToken", newAccessToken, 3600, "/", domain, false, true)
}

// rate limiting
func RateLimitMiddleware(c *gin.Context) {

}

// cors handling
func CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Next()
}

//
