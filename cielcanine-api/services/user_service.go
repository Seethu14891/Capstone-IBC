// cielcanine-api/services/user_service.go
package services

import (
	"context"

	"cielcanine-api/blockchain"
	"cielcanine-api/models"
	database "cielcanine-api/repositories"
)

func RegisterUser(user models.User) error {
	// Validate user data

	// Store user data in the database
	db := database.GetDB()
	_, err := db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	// Invoke chaincode for user registration
	err = blockchain.InvokeChaincodeForUserRegistration()
	if err != nil {
		return err
	}

	return nil
}
