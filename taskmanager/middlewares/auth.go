package middlewares

import (
	"net/http"

	"taskmanager/utils"

	"github.com/gin-gonic/gin"
)

// checks for valid JWT in cookies
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
