package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func createProductsTable(db *sql.DB) {
	statement := `CREATE TABLE products (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT,
		"price" REAL,
		"quantity" INTEGER
	);`

	log.Println("Creating products table")
	query, err := db.Prepare(statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	query.Exec()
	log.Println("Products table created")
}

func SetupSQLite() {
	// Create SQLite file if it doesn't exist
	if _, err := os.Stat("./data/sqlite.db"); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create("./data/sqlite.db")
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()

		// Connect to DB and create table
		connection := ConnectToDB()
		createProductsTable(connection)
		connection.Close()
	}
}
