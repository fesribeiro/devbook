package db

import (
	"database/sql"
	"devbook-api/src/config"

	_ "github.com/go-sql-driver/mysql"
)


func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StrDBConnection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}