package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// function to generate a JWT token for a user
func GenerateJWT(userId string) (string, error) {

	if userId == "" {
		return "", fmt.Errorf("User id is required")
	}
	// creat a claim
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // token expires in 72 hours
		"type":   "access",
	}
	// create a token with the claims and sign it with a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("jwt_secret")
	return token.SignedString([]byte(secret))

}

// refresh token

func RefreshJwt(userId string) (string, error) {

	if userId == "" {
		return "", fmt.Errorf("User id is required")
	}

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
		"type":   "refresh",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("jwt_secret")
	return token.SignedString([]byte(secret))

}
