package controllers

import (
	"devbook-api/src/auth"
	"devbook-api/src/db"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"devbook-api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userDB, err := repositories.NewUserRepository(db).FindByEmail(user.Email)

	if err != nil {
		fmt.Println(err, user)
		responses.Error(w, http.StatusUnauthorized, errors.New("user or password invalid"))
		return
	}

	if err := security.VerifyPassword(userDB.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("user or password invalid"))
		return
	}

	token, _ := auth.GenerateToken(userDB.ID)
	
	responses.JSON(w, http.StatusAccepted, map[string]string{
		"access_token": token,
	})
}