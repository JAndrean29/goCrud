package sqlitedb

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func InitiateSqliteConnection() (db *sqlx.DB) {
	db, err := sqlx.Connect("sqlite3", "./users.db")
	if err != nil {
		log.Fatal("db connection failed!")
	}

	return
}
