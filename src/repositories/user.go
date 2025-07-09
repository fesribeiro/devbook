package repositories

import (
	"database/sql"
	app_errors "devbook-api/src/errors"
	"devbook-api/src/models"
	"fmt"
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

func (userRepository user) Find(search string) ([]models.User, error) {
	db := userRepository.db
	defer db.Close()

	searchFormatted := fmt.Sprintf("%%%s%%", search)

	usersDB, err := db.Query("SELECT * FROM users WHERE name like ? or nick like ?", searchFormatted, searchFormatted)

	var users []models.User
	
	if err != nil {
		return nil, err
	}
	
	defer usersDB.Close()

	for usersDB.Next() {
		var user models.User

		if err := usersDB.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}


func (userRepository user) FindById(ID uint64) (*models.User, error) {
	db := userRepository.db
	defer db.Close()

	userDB, err := db.Query("SELECT * FROM users WHERE id = ?", ID)

	
	if err != nil {
		return nil, err
	}
	
	defer userDB.Close()

	var user models.User

	if userDB.Next() {

		if err := userDB.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}

	} else {
		return nil, app_errors.NewNotFoundError("User not found")
	}

	return &user, nil
}