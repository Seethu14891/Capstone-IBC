package handlers

import (
	"cielcanine-api/models"
	"cielcanine-api/utils"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Check user credentials (example: hardcoded for demonstration)
	if user.Username == "user" && user.Password == "password" {
		token, err := utils.GenerateToken(user.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Respond with token
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// Implementation for user signup (not included in this example)
}
