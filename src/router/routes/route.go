package routes

import (
	"devbook-api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI 	   string
	Method   string
	Func 	   func(http.ResponseWriter, *http.Request)
	WithAuth bool
}


func ConfigRoute(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoutes...)

	for _, route := range routes {

		if (route.WithAuth) {
			r.HandleFunc(route.URI, middlewares.Authenticated(route.Func)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Func).Methods(route.Method)
		}

	}

	return r
}