package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response JSON default
func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
	
}

// Response erro JSON default
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}

// Response erro JSON default
func CustomError(w http.ResponseWriter, statusCode int, err []string) {
	JSON(w, statusCode, struct {
		Err []string `json:"error"`
	}{
		Err: err,
	})
}