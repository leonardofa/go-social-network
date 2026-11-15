package controller

import (
	"net/http"
	"web/src/util"
)

// RenderPageSignup load SignUp page
func RenderPageSignup(w http.ResponseWriter, r *http.Request) {
	util.RenderTemplate(w, "signup.html", nil)
}
