package controller

import (
	"net/http"
	"web/src/util"
)

// RenderPageSignin load SignIn page
func RenderPageSignin(w http.ResponseWriter, r *http.Request) {
	util.RenderTemplate(w, "signin.html", nil)
}
