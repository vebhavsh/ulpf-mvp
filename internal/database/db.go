package database

import (
	"database/sql"
	"log"

	// This is a pure Go SQLite driver (doesn't require C compilers on Windows)
	_ "modernc.org/sqlite"
)

// DB is a global variable that holds the database connection
var DB *sql.DB

// InitDB creates the database file and sets up our table (Excel sheet equivalent)
func InitDB() {
	var err error

	// Open or create a file named 'ulpf.db' in your project folder
	DB, err = sql.Open("sqlite", "./ulpf.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// This SQL command creates our 'logs' table if it doesn't already exist.
	// It has 4 columns: id, timestamp, raw_log, and parsed_json.
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		raw_log TEXT,
		parsed_json TEXT
	);`

	// Execute the table creation command
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	log.Println("Database initialized successfully! Table 'logs' is ready.")
}

// InsertLog takes the original string and the final JSON, and saves them into a new row
func InsertLog(rawLog string, parsedJSON string) error {
	// The '?' are placeholders to prevent SQL injection (security best practice)
	insertSQL := `INSERT INTO logs(raw_log, parsed_json) VALUES (?, ?)`

	_, err := DB.Exec(insertSQL, rawLog, parsedJSON)
	if err != nil {
		return err
	}

	return nil
}
