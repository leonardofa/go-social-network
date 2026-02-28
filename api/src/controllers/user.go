package controllers

import (
	"net/http"
)

// CreateUser creates a new user.
func CreateUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("User created successfully"))
}

// ReadUser retrieves a single user by ID.
func ReadUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("User found successfully"))
}

// ReadUserList retrieves all users.
func ReadUserList(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("User(s) found successfully"))
}

// UpdateUser updates an existing user.
func UpdateUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("User updated successfully"))
}

// DeleteUser deletes a user by ID.
func DeleteUser(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("User deleted successfully"))
}
