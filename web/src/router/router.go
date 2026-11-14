package router

import (
	"web/src/router/path"

	"github.com/gorilla/mux"
)

// Init initializes and returns the WEB pages router with all configured paths.
func Init() *mux.Router {
	router := mux.NewRouter()

	return path.Init(router)
}
