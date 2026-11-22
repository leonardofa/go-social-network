// Package path defines user endpoints paths for user register.
package path

import (
	"net/http"
	"web/src/adapter"
)

// userEndpoints defines user endpoints.
var userEndpoints = []Path{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   adapter.CreateUser,
		Secure: false,
	},
}
