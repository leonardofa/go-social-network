package path

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Path helper representation.
type Path struct {
	// URI is the endpoint path.
	URI string
	// Method is the HTTP method.
	Method string
	// Func is the handler function.
	Func func(http.ResponseWriter, *http.Request)
	// Secure indicates if the endpoint requires authentication.
	Secure bool
}

// Init initializes the paths in router.
func Init(router *mux.Router) *mux.Router {
	paths := signin
	paths = append(paths, signup...)

	for _, path := range paths {
		_func := path.Func
		router.HandleFunc(path.URI, _func).Methods(path.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
