package main

import (
	controllers "cielcanine-api/controllers"
	"cielcanine-api/handlers"
	database "cielcanine-api/repositories"
	"log"
	"net/http"
)

func main() {

	// Initialize database
	database.Connect()

	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/signup", handlers.Signup)
	http.HandleFunc("/register", controllers.RegisterUserHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
