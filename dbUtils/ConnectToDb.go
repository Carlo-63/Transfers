package dbUtils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
