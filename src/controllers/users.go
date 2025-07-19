package controllers

import (
	"devbook-api/src/db"
	app_errors "devbook-api/src/errors"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	strErros, err := user.Validate("store")
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

	responses.JSON(w, http.StatusCreated, user)
}

func Find(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.URL.Query().Get("search"))

	db, err := db.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	users, err := repositories.NewUserRepository(db).Find(search)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func FindByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	user, err := repositories.NewUserRepository(db).FindById(userID)

	if err != nil {
		var notFoundErr *app_errors.NotFoundError
		if errors.As(err, &notFoundErr) {
			app_errors.NewNotFoundError("User not found").HttpError(w)
			return
		}

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func Update(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

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

	strErros, err := user.Validate("update")
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
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	defer db.Close()

	if err := repositories.NewUserRepository(db).Update(userID, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := repositories.NewUserRepository(db).Delete(userID); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
