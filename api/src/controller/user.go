package controller

import (
	"api/src/auth"
	"api/src/config"
	"api/src/model"
	"api/src/repository"
	"api/src/response"

	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUser creates a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare("create"); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	user.ID, err = repository.New(dbConn).Create(user)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

// ReadUserList retrieves all users.
func ReadUserList(w http.ResponseWriter, r *http.Request) {
	nameOrNick := r.URL.Query().Get("user")

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	users, err := repository.New(dbConn).FindByNameOrNick(nameOrNick)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

// ReadUser retrieves a single user by ID.
func ReadUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	user, err := repository.New(dbConn).FindByID(userID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID == 0 {
		response.ERROR(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	response.JSON(w, http.StatusOK, user)

}

// UpdateUser updates an existing user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := auth.ExtractUserIDFromRequest(r)
	switch {
	case err != nil:
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	case userID != userIDToken:
		response.ERROR(w, http.StatusForbidden, errors.New("you are not allowed to delete this user"))
		return
	}

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

	if err = user.Prepare("update"); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	err = repository.New(dbConn).Update(userID, user)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes a user by ID.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := auth.ExtractUserIDFromRequest(r)
	switch {
	case err != nil:
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	case userID != userIDToken:
		response.ERROR(w, http.StatusForbidden, errors.New("you are not allowed to delete this user"))
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	err = repository.New(dbConn).DeleteByID(userID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// FollowUser handles the user follow request by extracting the user ID from the request and performing the follow operation.
func FollowUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	followerID, err := auth.ExtractUserIDFromRequest(r)
	switch {
	case err != nil:
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	case followerID == userID:
		response.ERROR(w, http.StatusBadRequest, errors.New("you cannot follow yourself"))
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	err = repository.New(dbConn).Follow(userID, followerID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser removes the follower relationship between the authenticated user and the specified user.
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	followerID, err := auth.ExtractUserIDFromRequest(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	err = repository.New(dbConn).Unfollow(userID, followerID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
