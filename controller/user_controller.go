package controller

import (
	"e-commerce/model"
	"e-commerce/service"
	"e-commerce/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

// NewUserController returns a new instance of UserController
func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// GetUsers handles GET /users - returns all users
func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		utils.HandleError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	if users == nil {
		users = []model.User{} // Return empty array instead of null
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID handles GET /users/{id} - returns a single user by ID
func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := c.userService.GetUserByID(id)
	if err != nil {
		if err.Error() == "user not found" {
			utils.HandleError(w, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error fetching user %d: %v", id, err)
		utils.HandleError(w, http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser handles POST /users - creates a new user
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if err := validateUser(&user); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create user
	id, err := c.userService.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		utils.HandleError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Return success response with location header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"status":  "success",
		"message": "User created successfully",
		"data": map[string]interface{}{
			"id": id,
		},
	}
	json.NewEncoder(w).Encode(response)
}

// UpdateUser handles PUT /users/{id} - updates an existing user
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL path
	id, err := utils.ParseID(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Decode request body
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Set the ID from URL
	user.ID = id

	// Validate required fields (at least some fields should be present)
	if user.Name == "" && user.Email == "" && user.Password == "" {
		utils.HandleError(w, http.StatusBadRequest, "At least one field must be provided for update")
		return
	}

	// Update user
	rowsAffected, err := c.userService.UpdateUser(user)
	if err != nil {
		log.Printf("Error updating user %d: %v", id, err)
		if err.Error() == "user not found" {
			utils.HandleError(w, http.StatusNotFound, "User not found")
			return
		}
		utils.HandleError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	if rowsAffected == 0 {
		utils.HandleError(w, http.StatusNotFound, "User not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "User updated successfully",
	})
}

// DeleteUser handles DELETE /users/{id} - deletes a user by ID
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL path
	id, err := utils.ParseID(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Delete user
	rowsAffected, err := c.userService.DeleteUser(id)
	if err != nil {
		log.Printf("Error deleting user %d: %v", id, err)
		if err.Error() == "user not found" {
			utils.HandleError(w, http.StatusNotFound, "User not found")
			return
		}
		utils.HandleError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	if rowsAffected == 0 {
		utils.HandleError(w, http.StatusNotFound, "User not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// validateUser validates required user fields
func validateUser(user *model.User) error {
	if user.Name == "" {
		return fmt.Errorf("name is required")
	}
	if user.Email == "" {
		return fmt.Errorf("email is required")
	}
	if user.Password == "" {
		return fmt.Errorf("password is required")
	}
	if len(user.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}
	return nil
}