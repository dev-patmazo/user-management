package controllers

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/dev-patmazo/user-management/models"
	"github.com/gorilla/mux"
)

// CreateUser creates a new user based on the provided request body.
// It expects a JSON object containing the user details, including
// the username, age, role, and email.
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if newUser["username"] == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	if newUser["age"].(float64) < 0 {
		http.Error(w, "Age must be a positive number", http.StatusBadRequest)
		return
	}

	if newUser["role"] == "" {
		http.Error(w, "Role is required", http.StatusBadRequest)
		return
	}

	_, err = mail.ParseAddress(newUser["email"].(string))
	if err != nil {
		http.Error(w, "Invalid email address", http.StatusBadRequest)
		return
	}

	result, err := models.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "User created successfully",
		"data": map[string]interface{}{
			"username":   result.Username,
			"email":      result.Email,
			"created_at": result.CreatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}

}

// GetUser retrieves a user by their ID and returns the user
// information as a JSON response.
func GetUser(w http.ResponseWriter, r *http.Request) {

	requestParams := mux.Vars(r)
	id := requestParams["id"]

	result, err := models.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "User retrieved successfully",
		"data": map[string]interface{}{
			"id":       result.ID,
			"username": result.Username,
			"email":    result.Email,
			"age":      result.Age,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}
}

// UpdateUser updates an existing user with the provided data.
// It expects a JSON payload containing the user data in the request body.
// If the user data is valid, it calls the `UpdateUser` function from the `models` package to update the user in the database.
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	requestParams := mux.Vars(r)
	id := requestParams["id"]

	var existingUser map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&existingUser)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	if existingUser["username"] == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	if len(existingUser["password"].(string)) < 8 {
		http.Error(w, "Password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	_, err = mail.ParseAddress(existingUser["email"].(string))
	if err != nil {
		http.Error(w, "Invalid email address", http.StatusBadRequest)
		return
	}

	result, err := models.UpdateUser(id, existingUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "User updated successfully",
		"data": map[string]interface{}{
			"username":   result.Username,
			"email":      result.Email,
			"updated_at": result.UpdatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}

}

// DeleteUser deletes a user with the specified ID.
// The user ID is extracted from the request parameters and used to delete the user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	requestParams := mux.Vars(r)
	id := requestParams["id"]

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	err := models.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "User deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}
}
