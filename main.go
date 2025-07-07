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

	fmt.Println(config.ApiPort, config.StrDBConnection)

	fmt.Println("Starting server on port 5000...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), router.GenerateRouter()))
}
