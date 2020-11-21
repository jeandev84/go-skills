package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS book
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	author TEXT,
	page_count INTEGER
)
`

// Book Struct
type Book struct {
	Id        int
	Title     string
	Author    string
	pageCount int
}

// _ (underscore siginit le package est la mais je ne vais pas l'utiliser explicitement)
// Lancer : $ go build && ./goweb
func main() {

	// Ouverture de notre base de donnees
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	// Fermer la connection
	defer db.Close()

	// Ping connection
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	// Prepare schema
	stmt, err := db.Prepare(schema)

	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()

	// map to values
	var id int
	var title string
	var author string
	var pageCount int

	/*
	 * Method 1 : Recuperation de donnees
	 */
	rows, _ := db.Query("SELECT * FROM book")

	for rows.Next() {

		// On scanne chacune des colonnes en resultat
		rows.Scan(&id, &title, &author, &pageCount)
		fmt.Printf("id=%v, title=%v, author=%v, pageCount=%v\n",
			id,
			title,
			author,
			pageCount)
	}

	/*
	 * Method 2 : Recuperation des donnees
	 * map to struct
	 */
	rows, _ = db.Query("SELECT * FROM book")

	b := Book{}

	for rows.Next() {
		rows.Scan(&b.Id, &b.Title, &b.Author, &b.pageCount)
		fmt.Printf("book=%v\n", b)
	}

	// Si la connection a ete etablie alors il s'affichera le message suivant
	fmt.Println("Ping to database sussessful!")
}
