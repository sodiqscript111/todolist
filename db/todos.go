package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite" // Anonymous import: registers sqlite driver with database/sql
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite", "./todos.db")
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Enable foreign key support
	_, err = DB.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		log.Fatalf("Failed to enable foreign key support: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Create tables
	err = createUserTable()
	if err != nil {
		log.Fatalf("Could not create user table: %v", err)
	}

	err = createTable()
	if err != nil {
		log.Fatalf("Could not create todos table: %v", err)
	}
}

func createTable() error {
	createTodos := `
	CREATE TABLE IF NOT EXISTS todos (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		TITLE TEXT,
		DESCRIPTION TEXT,
		COMPLETED BOOLEAN,
		USER_ID INTEGER,
		FOREIGN KEY(USER_ID) REFERENCES users(ID)
	)`
	_, err := DB.Exec(createTodos)
	return err
}

func createUserTable() error {
	createUserTable := `
CREATE TABLE IF NOT EXISTS users (
    EMAIL TEXT,
    PASSWORD TEXT,
    ID INTEGER PRIMARY KEY AUTOINCREMENT
)`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		return err
	}
	return nil
}
