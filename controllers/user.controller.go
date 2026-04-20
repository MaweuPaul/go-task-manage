package controllers

// This file contains the handlers for user-related operations such as creating, updating, deleting, and retrieving users.

import (
	"taskmanager/models"
	"taskmanager/services"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler handles the creation of a new user.

func CreateUserHandler(c *gin.Context, email string, password string, nameFirst string, nameLast string, role string) {

	// Bind the JSON payload to a User struct
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Check if email and password are provided
	if email == "" || password == "" {
		c.JSON(400, gin.H{"error": "Email and password are required"})
		return
	}

	//  function to check if email is already in use

	// chekc if the email is already in use
	if services.IsEmailInUse(email) {
		c.JSON(400, gin.H{"error": "Email is already in use"})
		return
	}
}
