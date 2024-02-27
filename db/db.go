package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("sqlite3", "api.DB")
	if err != nil {
		panic("Error connecting to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTable()
}

func CreateTable() {
	createUserTable := `
  CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
  )
  `
	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Error creating users table.")
	}

	createEventsTable := `
  CREATE TABLE IF NOT EXISTS events (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  location TEXT NOT NULL,
  dateTime DATETIME NOT NULL,
  user_id INTEGER
  )
  `
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Error creating events table.")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	event_id INTEGER NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (event_id) REFERENCES events(id)
	)`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("Error creating registrations table.")
	}
}
