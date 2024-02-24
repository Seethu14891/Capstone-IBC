// utils/jwt_secret.go

package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateJWTSecretKey generates a JWT secret key of the specified length.
func GenerateJWTSecretKey(length int) (string, error) {
	// Calculate the byte size required for the given length
	byteSize := length / 4 * 3

	// Generate random bytes
	bytes := make([]byte, byteSize)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Encode the random bytes to base64
	encodedKey := base64.URLEncoding.EncodeToString(bytes)

	// Trim padding characters from the encoded key
	encodedKey = encodedKey[:length]

	return encodedKey, nil
}
