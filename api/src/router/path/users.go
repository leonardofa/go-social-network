// Package path defines user endpoints paths for user.
package path

import (
	"net/http"

	"api/src/controllers"
)

// UserPath defines all user-related API endpoints.
var userPath = []Path{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controllers.CreateUser,
		Secure: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodGet,
		Func:   controllers.ReadUser,
		Secure: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func:   controllers.ReadUserList,
		Secure: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodPut,
		Func:   controllers.UpdateUser,
		Secure: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodDelete,
		Func:   controllers.DeleteUser,
		Secure: false,
	},
}
