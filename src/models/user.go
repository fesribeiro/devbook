package models

import (
	"devbook-api/src/security"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty" validate:"required,min=5,max=50"`
	Nick      string    `json:"nick,omitempty" validate:"required"`
	Email     string    `json:"email,omitempty" validate:"required,email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *User) Validate(step string) ([]string, error) {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		errors := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field '%s' is not valid: %s", strings.ToLower(err.Field()), err.Tag()))
		}
		return errors, err
	}

	if err := u.format(step); err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if step == "store" {
		passwordHash, err := security.Hash(u.Password)

		if err != nil {
			return err
		}

		u.Password = string(passwordHash)
	}

	return nil
}
