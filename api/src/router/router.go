// Package routers manager API path routers.
package router

import (
	"api/src/router/path"

	"github.com/gorilla/mux"
)

// Execute initializes and returns the API router with all configured paths.
func Execute() *mux.Router {
	router := mux.NewRouter()

	return path.Init(router)
}
