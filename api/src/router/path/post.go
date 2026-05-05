package path

import (
	"api/src/controller"
	"net/http"
)

// postPaths define all post-related API endpoints.
var postPaths = []Path{
	{
		URI:    "/posts",
		Method: http.MethodPost,
		Func:   controller.CreatePost,
		Secure: true,
	},
	{
		URI:    "/posts",
		Method: http.MethodGet,
		Func:   controller.ReadPosts,
		Secure: true,
	},
	{
		URI:    "/posts/{id}",
		Method: http.MethodGet,
		Func:   controller.ReadPost,
		Secure: true,
	},
	{
		URI:    "/posts/{id}",
		Method: http.MethodPut,
		Func:   controller.UpdatePost,
		Secure: true,
	},
	{
		URI:    "/posts/{id}",
		Method: http.MethodDelete,
		Func:   controller.DeletePost,
		Secure: true,
	},
	{
		URI:    "/users/{id}/posts",
		Method: http.MethodGet,
		Func:   controller.ReadUserPosts,
		Secure: true,
	},
	{
		URI:    "/posts/{id}/like",
		Method: http.MethodPost,
		Func:   controller.LikePost,
		Secure: true,
	},
	{
		URI:    "/posts/{id}/unlike",
		Method: http.MethodPost,
		Func:   controller.UnlikePost,
		Secure: true,
	},
}
