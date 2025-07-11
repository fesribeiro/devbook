package middlewares

import (
	"fmt"
	"net/http"
)

func Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Auth...")

		next(w, r)
	}
}