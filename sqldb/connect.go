package sqldb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "events.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		venue TEXT NOT NULL,
		created_by TEXT NOT NULL,
		date_time DATETIME NOT NULL
	);
	`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}
