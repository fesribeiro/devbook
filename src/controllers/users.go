package userController

import "net/http"

func Store(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creating user"))
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