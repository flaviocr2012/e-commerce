package controller

import (
	"e-commerce/model"
	"e-commerce/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get the 'id' from the query string
	idStr := r.URL.Query().Get("id")

	// Convert the 'id' from string to int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the service method with the converted id
	user, err := c.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Encode the user to JSON and write to the response
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a User struct
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service method with the user struct
	id, err := c.userService.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the ID of the created user to the response
	json.NewEncoder(w).Encode(id)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a User struct
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service method with the user struct
	rowsAffected, err := c.userService.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the number of rows affected to the response
	json.NewEncoder(w).Encode(rowsAffected)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get the 'id' from the query string
	idStr := r.URL.Query().Get("id")

	// Convert the 'id' from string to int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the service method with the converted id
	rowsAffected, err := c.userService.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the number of rows affected to the response
	json.NewEncoder(w).Encode(rowsAffected)
}
