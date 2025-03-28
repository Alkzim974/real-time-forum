package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "./db/forum.db")
	if err != nil {
		log.Fatal(err)
	}

	CreateUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		nickname TEXT UNIQUE NOT NULL,
		age INTEGER NOT NULL,
		gender TEXT CHECK( gender IN ('Male', 'Female', 'Other') ) NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password BLOB NOT NULL
	);`

	_, err = DB.Exec(CreateUserTable)
	if err != nil {
		log.Fatal(err)
	}

}
