package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)


func main () {
	config.Load()

	fmt.Printf("Starting server on port %d...\n", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), router.GenerateRouter()))
}
