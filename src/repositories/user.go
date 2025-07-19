package repositories

import (
	"database/sql"
	app_errors "devbook-api/src/errors"
	"devbook-api/src/models"
	"fmt"
	"strings"
)

type user struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

func (userRepository user) Store(user models.User) (uint64, error) {
	db := userRepository.db

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

	searchFormatted := fmt.Sprintf("%%%s%%", search)

	usersDB, err := db.Query("SELECT * FROM users WHERE name like ? or nick like ?", searchFormatted, searchFormatted)

	users := []models.User{}

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

func (userRepository user) Update(ID uint64, userData *models.User) error {
	exists, err := userRepository.exists(ID, userData)

	if err != nil {
		return err
	}

	if exists != nil {
		return app_errors.NewNotFoundError(strings.Join(exists, ";"))
	}

	db := userRepository.db

	statement, err := db.Prepare("UPDATE users set name = ?, nick = ?, email = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(userData.Name, userData.Nick, userData.Email, ID); err != nil {
		return err
	}

	return nil
}

func (userRepository user) exists(ID uint64, userData *models.User) ([]string, error) {

	var user models.User

	errChan := make(chan string)

	db := userRepository.db

	userDB, err := db.Query("SELECT nick, email FROM users WHERE id = ?", ID)

	if err != nil {
		return nil, err
	}

	defer userDB.Close()

	if !userDB.Next() {
		return []string{"User do not exists"}, nil
	}

	if err = userDB.Scan(&user.Nick, &user.Email); err != nil {
		return nil, err
	}

	go func() {
		nickNameExists, err := db.Query("SELECT nick FROM users WHERE nick = ? AND id != ?", userData.Nick, ID)

		if err != nil {
			errChan <- err.Error()
			return
		}

		defer nickNameExists.Close()

		if nickNameExists.Next() {
			errChan <- "nickname already exists"
			return
		}

		errChan <- ""
	}()

	go func() {
		emailExists, err := db.Query("SELECT email FROM users WHERE email = ? AND id != ? LIMIT 1", userData.Email, ID)

		if err != nil {
			errChan <- err.Error()
			return
		}

		defer emailExists.Close()

		if emailExists.Next() {
			errChan <- "email already exists"
			return
		}

		errChan <- ""
	}()

	var queryErrors []string

	for range 2 {
		msg := <-errChan
		if msg != "" {
			queryErrors = append(queryErrors, msg)
		}
	}

	return queryErrors, nil
}

func (userRepository user) Delete(ID uint64) error {
	db := userRepository.db

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (userRepository user) FindByEmail(email string) (*models.User, error) {
	db := userRepository.db

	userDB, err := db.Query("SELECT * FROM users WHERE email = ?", email)

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
