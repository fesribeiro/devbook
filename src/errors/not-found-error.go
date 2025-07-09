package app_errors

import (
	"devbook-api/src/responses"
	"net/http"
)

type NotFoundError struct {
	Msg string
}

func (e *NotFoundError) Error() string {
	return e.Msg
}

func (e *NotFoundError) HttpError(w http.ResponseWriter) {
	responses.Error(w, http.StatusNotFound, e)
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{Msg: msg}
}