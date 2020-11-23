package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

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

	fmt.Println("Ping to database successful!")
}
