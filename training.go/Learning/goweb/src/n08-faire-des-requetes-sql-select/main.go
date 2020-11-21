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
)
`

type Book struct {
	Id     int
	Title  string
	Author string
}

func main() {
	// Open connection
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Printf("Cannot open database. err=%v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Create database
	stmt, err := db.Prepare(schema)
	if err != nil {
		fmt.Printf("Cannot execute SQL query. err=%v\n", err)
		os.Exit(1)
	}
	stmt.Exec()

	// map to values
	rows, _ := db.Query("SELECT * FROM book")
	var id int
	var title string
	var author string

	for rows.Next() {
		rows.Scan(&id, &title, &author)
		fmt.Printf("id=%v, title=%v, author=%v\n", id, title, author)
	}

	// map to struct
	rows, _ = db.Query("SELECT * FROM book")
	b := Book{}

	for rows.Next() {
		rows.Scan(&b.Id, &b.Title, &b.Author)
		fmt.Printf("book=%v\n", b)
	}

}
