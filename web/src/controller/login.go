package controller

import (
	"net/http"
)

// LoginPage load login page
func LoginPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Login Page"))
	if err != nil {
		return
	}
}
