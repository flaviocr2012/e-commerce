package main

import (
	"database/sql"
	"e-commerce/controller"
	"e-commerce/repository"
	"e-commerce/router"
	"e-commerce/service"
	"log"
	"net/http"
)

func main() {

	// Database connection
	connStr := "user=flavio password=securepassword dbname=flaviodb host=localhost port=5433 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize your repository
	userRepository := repository.NewUserRepository(db)

	// Initialize your services
	userService := service.NewUserService(userRepository)

	// Initialize your controllers
	userController := controller.NewUserController(userService)

	// Set up the router
	r := router.SetUpRouter(userController)

	// Start the server
	http.ListenAndServe(":8080", r)
}
