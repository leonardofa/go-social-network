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

// CreatePost handles the creation of a new post by the authenticated user and returns the created post in the response.
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserIDFromRequest(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post model.Post
	if err = json.Unmarshal(body, &post); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare(); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := config.GetConnection()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	repo := repository.NewPostsRepository(dbConn)
	post.AuthorID = userID
	post.ID, err = repo.Create(post)
	post.CreatedAt = nil
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, post)
}

// ReadPosts busca os posts que apareceriam no feed do usuário.
func ReadPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserIDFromRequest(r)
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

	repo := repository.NewPostsRepository(dbConn)
	posts, err := repo.FindAll(userID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

// ReadPost handles an HTTP request to retrieve a post by its ID from the database and sends the post as a JSON response.
func ReadPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 64)
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

	repo := repository.NewPostsRepository(dbConn)
	post, err := repo.FindByID(postID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if post.ID == 0 {
		response.ERROR(w, http.StatusNotFound, errors.New("post not found"))
		return
	}

	response.JSON(w, http.StatusOK, post)
}

// UpdatePost altera os dados de um post.
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserIDFromRequest(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 64)
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

	repo := repository.NewPostsRepository(dbConn)
	postInDB, err := repo.FindByID(postID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if postInDB.ID == 0 {
		response.ERROR(w, http.StatusNotFound, errors.New("post not found"))
		return
	}

	if postInDB.AuthorID != userID {
		response.ERROR(w, http.StatusForbidden, errors.New("you cannot update a post that is not yours"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post model.Post
	if err = json.Unmarshal(body, &post); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare(); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	post.ID = postID
	if err = repo.Update(postID, post); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// DeletePost exclui os dados de um post.
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserIDFromRequest(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 64)
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

	repo := repository.NewPostsRepository(dbConn)
	postInDB, err := repo.FindByID(postID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if postInDB.ID == 0 {
		response.ERROR(w, http.StatusNotFound, errors.New("post not found"))
		return
	}

	if postInDB.AuthorID != userID {
		response.ERROR(w, http.StatusForbidden, errors.New("you cannot delete a post that is not yours"))
		return
	}

	if err = repo.Delete(postID); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// ReadUserPosts busca todos os posts de um usuário.
func ReadUserPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
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

	repo := repository.NewPostsRepository(dbConn)
	posts, err := repo.FindByUserID(userID)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}

// LikePost adiciona uma curtida no post.
func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 64)
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

	repo := repository.NewPostsRepository(dbConn)
	if err = repo.Like(postID); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// UnlikePost subtrai uma curtida no post.
func UnlikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 64)
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

	repo := repository.NewPostsRepository(dbConn)
	if err = repo.Unlike(postID); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
