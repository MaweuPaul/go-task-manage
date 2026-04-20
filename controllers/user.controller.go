package controllers

import (
	"os"
	"taskmanager/models"
	"taskmanager/services"

	"github.com/gin-gonic/gin"
)

func setTokenCookies(c *gin.Context, accessToken string, refreshToken string) {
	domain := os.Getenv("DOMAIN")
	secure := os.Getenv("SECURE") == "true"
	c.SetCookie("access_token", accessToken, 3600, "/", domain, secure, true)
	c.SetCookie("refresh_token", refreshToken, 7*24*3600, "/", domain, secure, true)
}

func CreateUserHandler(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := services.CreateUser(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	setTokenCookies(c, createdUser.AccessToken, createdUser.RefreshToken)
	c.JSON(201, gin.H{"message": "User created successfully", "user": createdUser.User})
}

func LoginUserHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := services.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	setTokenCookies(c, response.AccessToken, response.RefreshToken)
	c.JSON(200, gin.H{"message": "User logged in successfully", "user": response.User})
}
