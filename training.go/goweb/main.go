package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// _ (underscore siginit le package est la mais je ne vais pas l'utiliser explicitement)
func main() {

	// Ouverture de notre base de donnees
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	// Fermer la connection
	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ping to database sussessful!")
}
