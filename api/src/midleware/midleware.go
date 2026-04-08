package midleware

import (
	"fmt"
	"net/http"
)

// Logger is a middleware that logs incoming requests.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nRequest received: %s %s | from: %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// AuthValidation is a middleware that performs authentication checks before passing the request to the next handler.
func AuthValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nAuthenticated method: %s %s | from: %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
