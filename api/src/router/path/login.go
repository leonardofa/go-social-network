// Package path defines user endpoints paths for user.
package path

import (
	"api/src/controller"
	"net/http"
)

// loginPath defines the "/login" endpoint with HTTP POST method and an unsecured handler for login functionality.
var loginPath = Path{
	URI:    "/login",
	Method: http.MethodPost,
	Func:   controller.Login,
	Secure: false,
}
