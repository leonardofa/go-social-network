package controller

import (
	"api/src/config"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
)

// Login handles user authentication by processing the login request and validating user credentials.
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	match, err := repository.New(dbConn).Login(user)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if !match {
		response.JSON(w, http.StatusUnauthorized, nil)
		return
	}

	response.JSON(w, http.StatusOK, "Authenticated")
}
