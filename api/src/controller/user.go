package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"api/src/config"
	"api/src/model"
	"api/src/repository"
)

// CreateUser creates a new user.
func CreateUser(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Print(err)
		return
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Print(err)
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		log.Print(err)
		return
	}

	id, err := repository.New(dbConn).Create(user)
	if err != nil {
		log.Print(err)
		return
	}

	_, _ = writer.Write([]byte(fmt.Sprintf("User created successfully with ID: %d", id)))
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
