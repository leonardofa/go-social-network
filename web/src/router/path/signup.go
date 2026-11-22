// Package path defines user endpoints paths for user register.
package path

import (
	"net/http"
	"web/src/controller"
)

// signup defines signup paths.
var signup = []Path{
	{
		URI:    "/signup",
		Method: http.MethodGet,
		Func:   controller.RenderPageSignup,
		Secure: false,
	},
}
