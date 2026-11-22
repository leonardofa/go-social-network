// Package path defines user endpoints user authentication (sign in).
package path

import (
	"net/http"
	"web/src/controller"
)

// signin defines signin paths.
var signin = []Path{
	{
		URI:    "/",
		Method: http.MethodGet,
		Func:   controller.RenderPageSignin,
		Secure: false,
	},
	{
		URI:    "/signin",
		Method: http.MethodGet,
		Func:   controller.RenderPageSignin,
		Secure: false,
	},
}
