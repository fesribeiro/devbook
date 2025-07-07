package main

import (
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)


func main () {
	fmt.Println("Starting server on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", router.GenerateRouter()))
}
