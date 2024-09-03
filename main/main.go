package main

import (
	"e-commerce/controller"
	"e-commerce/router"
	"e-commerce/service"
	"net/http"
)

func main() {
	// Initialize your services
	userService := service.NewUserService()

	// Initialize your controllers
	userController := controller.NewUserController(userService)

	// Set up the router
	r := router.SetUpRouter(userController)

	// Start the server
	http.ListenAndServe(":8080", r)
}
