package routes

import (
	"devbook-api/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI: "/login",
		Method: http.MethodPost,
		Func: controllers.Login,
		WithAuth: false,
	},
}