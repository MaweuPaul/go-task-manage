// Package utils provides utility functions for the taskmanager application.
// This file contains JWT helper functions for generating and validating tokens.

package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a signed JWT access token for the given user ID.
// Parameters:
//   - userId: the unique identifier of the user
//
// Returns:
//   - string: the signed JWT access token
//   - error: if userId is empty or token generation fails
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
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))

}

// RefreshJwt generates a signed JWT refresh token for the given user ID.
// Parameters:
//   - userId: the unique identifier of the user
//
// Returns:
//   - string: the signed JWT refresh token valid for 30 days
//   - error: if userId is empty or token generation fails
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
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))

}

// ValidateToken parses and validates a JWT token string.
// Parameters:
//   - tokenString: the JWT token string to validate
//
// Returns:
//   - jwt.MapClaims: the decoded claims if token is valid
//   - error: if token is invalid, expired or parsing fails

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	if tokenString == "" {

		return jwt.MapClaims{}, fmt.Errorf("Token is required")
	}
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return jwt.MapClaims{}, err
	}
	if !token.Valid {
		return jwt.MapClaims{}, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
