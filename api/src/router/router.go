// Package routers manager API path routers.
package router

import (
	"api/src/router/path"

	"github.com/gorilla/mux"
)

// Init initializes and returns the API router with all configured paths.
func Init() *mux.Router {
	router := mux.NewRouter()

	return path.Init(router)
}
