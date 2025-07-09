package userController

import (
	"devbook-api/src/db"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Store(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	strErros, err := user.Validate();
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, struct {
			Error []string `json:"error"`
		}{
			Error: strErros,
		})
		return
	}

	db, err := db.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	lastIDCreated, err := repositories.NewUserRepository(db).Store(user)
	user.ID = lastIDCreated

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("User created with ID:", user)
	responses.JSON(w, http.StatusCreated, user)
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