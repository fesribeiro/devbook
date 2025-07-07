package routes

import (
	userController "devbook-api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodPost,
		Func: userController.Store,
		WithAuth: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Func: userController.FindAll,
		WithAuth: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodGet,
		Func: userController.FindByID,
		WithAuth: false,
	},
	{
		URI: "/users",
		Method: http.MethodPut,
		Func: userController.Update,
		WithAuth: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodDelete,
		Func: userController.Delete,
		WithAuth: false,
	},
}