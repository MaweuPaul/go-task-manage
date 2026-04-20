package services

import (
	"fmt"
	"taskmanager/models"
	"taskmanager/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]models.User)

func IsEmailInUse(email string) bool {
	for _, u := range users {
		if u.Email == email {
			return true
		}
	}
	return false
}

func CreateUser(input models.CreateUserInput) (models.AuthResponse, error) {

	if input.Email == "" || input.Password == "" || input.NameFirst == "" || input.NameLast == "" {
		return models.AuthResponse{}, fmt.Errorf("all fields are required")
	}

	if IsEmailInUse(input.Email) {
		return models.AuthResponse{}, fmt.Errorf("unable to complete registration")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to hash password: %v", err)
	}

	// create user from input
	user := models.User{
		ID:        uuid.New().String(),
		NameFirst: input.NameFirst,
		NameLast:  input.NameLast,
		Email:     input.Email,
		Role:      input.Role,
		Password:  string(hashedPassword),
	}

	// store user
	users[user.ID] = user

	// clear password before returning
	user.Password = ""

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

func LoginUser(email, password string) (models.AuthResponse, error) {

	if email == "" || password == "" {
		return models.AuthResponse{}, fmt.Errorf("email and password are required")
	}

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

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("invalid email or password")
	}

	user.Password = ""

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
