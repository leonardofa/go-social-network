// Package path defines user endpoints paths for user.
package path

import (
	"net/http"
	"web/src/controller"
)

// loginPath defines the "/login" endpoint with HTTP POST method and an unsecured handler for login functionality.
var loginPath = []Path{
	{
		URI:    "/",
		Method: http.MethodGet,
		Func:   controller.LoginPage,
		Secure: false,
	},
	{
		URI:    "/login",
		Method: http.MethodGet,
		Func:   controller.LoginPage,
		Secure: false,
	},
}
