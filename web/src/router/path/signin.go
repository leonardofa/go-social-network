// Package path defines user endpoints paths for user.
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
