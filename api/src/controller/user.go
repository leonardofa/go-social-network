package controller

import (
	"api/src/auth"
	"api/src/config"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"api/src/security"

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
	followingUserId, err := strconv.ParseUint(mux.Vars(r)["followingUserId"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userIDToken, err := auth.ExtractUserIDFromRequest(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	switch {
	case userID == followingUserId:
		response.ERROR(w, http.StatusForbidden, errors.New("you cannot follow yourself"))
		return
	case userIDToken != userID:
		response.ERROR(w, http.StatusForbidden, errors.New("you are not allowed to follow this user"))
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	err = repository.New(dbConn).Follow(userID, followingUserId)
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
	followingUserId, err := strconv.ParseUint(mux.Vars(r)["followingUserId"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userIDToken, err := auth.ExtractUserIDFromRequest(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	switch {
	case userID == followingUserId:
		response.ERROR(w, http.StatusForbidden, errors.New("you cannot unfollow yourself"))
		return
	case userIDToken != userID:
		response.ERROR(w, http.StatusForbidden, errors.New("you are not allowed to unfollow this user"))
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	err = repository.New(dbConn).Unfollow(userID, followingUserId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// ReadFollowersList retrieves the list of followers for a specified user.
func ReadFollowersList(w http.ResponseWriter, r *http.Request) {
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

	followers, err := repository.New(dbConn).ReadFollowersList(userID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

// ReadFollowingList retrieves the list of users that the specified user is following.
func ReadFollowingList(w http.ResponseWriter, r *http.Request) {
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

	followers, err := repository.New(dbConn).ReadFollowingList(userID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)
}

// UpdatePassword updates the password of a user
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userIDToken, err := auth.ExtractUserIDFromRequest(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if userIDToken != userID {
		response.ERROR(w, http.StatusForbidden, errors.New("you are not allowed to change the password of this user"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var passwordChangeModel model.Password
	if err = json.Unmarshal(body, &passwordChangeModel); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	if passwordChangeModel.Actual == "" || passwordChangeModel.New == "" {
		response.ERROR(w, http.StatusBadRequest, errors.New("actual and new password fields are required"))
		return
	}
	if passwordChangeModel.New == passwordChangeModel.Actual {
		response.ERROR(w, http.StatusBadRequest, errors.New("new password cannot be the same as the old password"))
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	if password, err := repository.New(dbConn).GetPassWord(userID); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	} else if err = security.CompareHashAndPassword(password, passwordChangeModel.Actual); err != nil {
		response.ERROR(w, http.StatusForbidden, errors.New("actual password is incorrect"))
		return
	}

	if newHash, err := security.GenerateFromPassword(passwordChangeModel.New); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	} else {
		if err = repository.New(dbConn).UpdatePassword(userID, string(newHash)); err != nil {
			response.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	}

	response.JSON(w, http.StatusNoContent, nil)
}
