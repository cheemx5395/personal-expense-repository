package repository

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func OpenDB() (Queries, error) {
	db, err := sql.Open("sqlite", "expenses.db")
	if err != nil {
		return Queries{}, err
	}

	if err := db.Ping(); err != nil {
		return Queries{}, err
	}

	return *New(db), nil
}
