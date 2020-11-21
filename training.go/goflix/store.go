package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

/*
 Interface pour faire des requettes SQL
 - Open  (ouvrir la connection)
 - Close (fermer la connection)
*/
type Store interface {
	Open() error
	Close() error

	GetMovies() ([]*Movie, error)
}

type dbStore struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS movie
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT, 
	release_date TEXT,
	duration INTEGER,
	trailer_url TEXT
)
`

// implementation de la methode Open() de l'interface Store
func (store *dbStore) Open() error {

	db, err := sqlx.Connect("sqlite3", "goflix.db")

	if err != nil {
		return err
	}

	log.Println("Connected to DB")

	db.MustExec(schema)

	store.db = db

	return nil
}

// implementation de la methode Close() de l'interface Store
func (store *dbStore) Close() error {
	return store.db.Close()
}

func (store *dbStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")

	if err != nil {
		return movies, err
	}

	return movies, nil
}
