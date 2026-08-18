package router

import (
	"e-commerce/controller"
	"e-commerce/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRouter(
	userController *controller.UserController,
	productController *controller.ProductController,
	cartController *controller.CartController,
) *mux.Router {
	r := mux.NewRouter()

	// Apply middleware
	r.Use(utils.JsonContentTypeMiddleware)

	// Root route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":    "API is running",
			"version":   "1.0",
			"endpoints": "/users, /users/{id}, /products, /products/{id}, /carts",
		})
	}).Methods("GET")

	// User routes
	r.HandleFunc("/users", userController.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Product routes
	r.HandleFunc("/products", productController.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", productController.GetProductByID).Methods("GET")
	r.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productController.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productController.DeleteProduct).Methods("DELETE")

	// Cart routes
	r.HandleFunc("/carts", cartController.GetCart).Methods("GET")
	r.HandleFunc("/carts/items", cartController.AddToCart).Methods("POST")
	r.HandleFunc("/carts/items/{id}", cartController.UpdateCartItem).Methods("PUT")
	r.HandleFunc("/carts/items/{id}", cartController.RemoveFromCart).Methods("DELETE")
	r.HandleFunc("/carts/clear", cartController.ClearCart).Methods("DELETE")

	return r
}