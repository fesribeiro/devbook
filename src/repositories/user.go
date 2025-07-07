package repositories

import (
	"database/sql"
	"devbook-api/src/models"
)

type user struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

func (userRepository user) Store(user models.User) (uint64, error) {
	db := userRepository.db
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?,?,?,?)")

	if err != nil {
		return 0, err
	}
	
	defer statement.Close()


	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastIDCreated, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastIDCreated), nil
}