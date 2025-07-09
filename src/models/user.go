package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID    		uint64 		`json:"id,omitempty"`
	Name  		string 		`json:"name,omitempty" validate:"required,min=5,max=50"`
	Nick  		string 		`json:"nick,omitempty" validate:"required"`
	Email 		string 		`json:"email,omitempty" validate:"required,email"`
	Password  string 		`json:"password,omitempty" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}


func (u User) Validate() ([]string, error) {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		errors := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field '%s' is not valid: %s", strings.ToLower(err.Field()), err.Tag()))
		}
		return errors, err
	}

	u.format()
	return nil, nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}