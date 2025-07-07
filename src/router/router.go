package router

import (
	"devbook-api/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	router := mux.NewRouter()
	return routes.ConfigRoute(router)
}