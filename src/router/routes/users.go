package routes

import (
	userController "devbook-api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Func:     userController.Store,
		WithAuth: true,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Func:     userController.Find,
		WithAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Func:     userController.FindByID,
		WithAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodPut,
		Func:     userController.Update,
		WithAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodDelete,
		Func:     userController.Delete,
		WithAuth: true,
	},
}
