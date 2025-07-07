package userController

import (
	"devbook-api/src/db"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Store(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	userID, err := repositories.NewUserRepository(db).Store(user)

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID Created: %d", userID)))
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("finding all user"))
}

func FindByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("finding user"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating user"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting user"))
}