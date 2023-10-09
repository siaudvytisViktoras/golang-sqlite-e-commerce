package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	// Initialize and open the SQLite database
	db, err := sql.Open("sqlite3", "your-database-file.db")
	if err != nil {
		return nil, err
	}

	// Create necessary tables if they don't exist
	// ...
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT,
			price REAL
		);
		`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
