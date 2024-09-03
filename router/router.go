package router

import (
	"e-commerce/controller"
	"e-commerce/utils"
	"github.com/gorilla/mux"
)

// SetUpRouter sets up the router with all the necessary routes and middlewares
func SetUpRouter(userController *controller.UserController) *mux.Router {
	r := mux.NewRouter()

	// Apply middleware
	r.Use(utils.JsonContentTypeMiddleware)

	// Define your routes
	r.HandleFunc("/users", userController.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	return r
}
