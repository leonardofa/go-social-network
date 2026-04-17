// Package path defines user endpoints paths for user.
package path

import (
	"net/http"

	"api/src/controller"
)

// UserPath defines all user-related API endpoints.
var userPath = []Path{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controller.CreateUser,
		Secure: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodGet,
		Func:   controller.ReadUser,
		Secure: true,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func:   controller.ReadUserList,
		Secure: true,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodPut,
		Func:   controller.UpdateUser,
		Secure: true,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
		Secure: true,
	},
}
