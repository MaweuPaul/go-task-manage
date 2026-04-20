package controllers

// This file contains the handlers for user-related operations such as creating, updating, deleting, and retrieving users.

import (
	"taskmanager/models"
	"taskmanager/services"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler handles the creation of a new user.

func CreateUserHandler(c *gin.Context) {

	// Bind the JSON payload to a User struct
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Create the user using the service layer
	createdUser, err := services.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, createdUser)
}
