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

// Init initializes the user path router.
func Init(router *mux.Router) *mux.Router {
	paths := userPath
	paths = append(paths, loginPath)

	for _, path := range paths {
		router.HandleFunc(path.URI, path.Func).Methods(path.Method)
	}

	return router
}
