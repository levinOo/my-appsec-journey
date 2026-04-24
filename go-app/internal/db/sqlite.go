package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB represents our database connection
var DB *sql.DB

// InitDB initializes an in-memory SQLite database and creates some test data
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		role TEXT
	);
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	_, _ = DB.Exec("INSERT INTO users (username, role) VALUES ('admin', 'admin')")
	_, _ = DB.Exec("INSERT INTO users (username, role) VALUES ('guest', 'user')")
}
