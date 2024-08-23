package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
	dbPath, err := initPath()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	initialiseQuery := `CREATE TABLE IF NOT EXISTS messages(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		message TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS groups(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		group_name TEXT NOT NULL
	);
	`
	_, err = db.Exec(initialiseQuery)
	if err != nil {
		return nil, err
	}
	return db, nil
}
