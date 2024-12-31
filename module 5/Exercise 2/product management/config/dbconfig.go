package db

import (
	"database/sql"
	"log"
	_"modernc.org/sqlite"
)

var DB *sql.DB

func InitialiseDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./Inventory.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			price REAL,
			quantity INTEGER,
			timestamp DATETIME
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return DB
}
