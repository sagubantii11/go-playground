package sqldb

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		emailID TEXT NOT NULL UNIQUE,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		full_name TEXT NOT NULL
	);
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		venue TEXT NOT NULL,
		created_by TEXT NOT NULL,
		userid INTEGER NOT NULL,
		date_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (created_by) REFERENCES users(username),
		FOREIGN KEY (userid) REFERENCES users(id)
	);
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}
