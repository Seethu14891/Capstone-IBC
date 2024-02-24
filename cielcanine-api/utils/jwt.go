package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generates a JWT token for the given username.
func GenerateToken(username string) (string, error) {
	// Generate JWT secret key
	secretKey, err := GenerateJWTSecretKey(32) // Specify the desired key length
	if err != nil {
		return "", err
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign the JWT token with the generated secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
