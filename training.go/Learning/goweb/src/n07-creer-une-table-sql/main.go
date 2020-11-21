package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS book
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    author TEXT
    page_count INTEGER
)
`

func main() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Printf("Cannot open database. err=%v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Printf("Cannot check database connection. err=%v\n", err)
		os.Exit(1)
	}

	stmt, err := db.Prepare(schema)
	if err != nil {
		fmt.Printf("Cannot execute SQL query. err=%v\n", err)
		os.Exit(1)
	}

	stmt.Exec()

	fmt.Println("Ping to database successful!")
}
