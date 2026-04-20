package services

import (
	"fmt"
	"taskmanager/models"
	"taskmanager/utils"

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

func CreateUser(user models.User) (models.AuthResponse, error) {

	// validate the user data
	if user.Email == "" || user.Password == "" || user.NameFirst == "" || user.NameLast == "" {
		return models.AuthResponse{}, fmt.Errorf("all fields are required")
	}

	// check if email is already in use
	if IsEmailInUse(user.Email) {
		return models.AuthResponse{}, fmt.Errorf("email is already in use")
	}

	// hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = string(hashedPassword)
	// id
	user.ID = uuid.New().String()
	// store the user in the in-memory map
	users[user.ID] = user

	// generate JWT tokens
	accessToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate access token: %v", err)
	}
	refreshToken, err := utils.RefreshJwt(user.ID)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	return models.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// login user with email and password

func LoginUser(email, password string) (models.AuthResponse, error) {

	if email == "" || password == "" {

		return models.AuthResponse{}, fmt.Errorf("email and password are required")
	}

	// find the user by email
	var user models.User
	found := false
	for _, u := range users {
		if u.Email == email {
			user = u
			found = true
			break
		}

	}

	if !found {
		return models.AuthResponse{}, fmt.Errorf("invalid email or password")
	}

	// verify the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("invalid email or password")
	}

	// generate JWT tokens
	accessToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate access token: %v", err)
	}
	refreshToken, err := utils.RefreshJwt(user.ID)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate refresh token: %v", err)
	}
	return models.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
