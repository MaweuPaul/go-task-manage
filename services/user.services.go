package services

import (
	"fmt"
	"taskmanager/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// in memory user storage
var users = make(map[string]models.User)

// function to check if email is already in use
func IsEmailInUse(email string) bool {
	// This is a placeholder implementation. In a real application,check the database for existing email.
	for _, u := range users {

		if u.Email == email {
			return true
		}
	}

	return false

}

// create a new user

func CreateUser(user models.User) (models.User, error) {

	// validate the user data
	if user.Email == "" || user.Password == "" || user.NameFirst == "" || user.NameLast == "" {
		return models.User{}, fmt.Errorf("all fields are required")
	}

	// check if email is already in use
	if IsEmailInUse(user.Email) {
		return models.User{}, fmt.Errorf("email is already in use")
	}

	// hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = string(hashedPassword)
	// id
	user.ID = uuid.New().String()
	// store the user in the in-memory map
	users[user.ID] = user

	return user, nil
}
