package controller

import (
	"e-commerce/model"
	"e-commerce/service"
	"e-commerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := c.userService.GetUserByID(id)
	if err != nil {
		utils.HandleError(w, http.StatusNotFound, "User not found")
		return
	}

	response := utils.Response{
		Status: "success",
		Data:   user,
	}
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	id, err := c.userService.CreateUser(user)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Set Location header
	w.Header().Set("Location", fmt.Sprintf("/users/%d", id))
	w.WriteHeader(http.StatusCreated)

	response := utils.Response{
		Status: "success",
		Data:   id,
	}
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a User struct
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service method with the user struct
	_, err = c.userService.UpdateUser(user)
	if err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return 200 OK for successful update
	w.WriteHeader(http.StatusOK)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get the 'id' from the query string
	idStr := r.URL.Query().Get("id")

	// Convert the 'id' from string to int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service method with the converted id
	_, err = c.userService.DeleteUser(id)
	if err != nil {
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return 204 No Content for successful deletion
	w.WriteHeader(http.StatusNoContent)
}
