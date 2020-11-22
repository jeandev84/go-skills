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

	// Connection
	Open() error
	Close() error

	// Movie methods
	GetMovies() ([]*Movie, error)
	GetMovieById(id int64) (*Movie, error)
	CreateMovie(m *Movie) error

	// User methods
	FindUser(username string, password string) (bool, error)
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
);

CREATE TABLE IF NOT EXISTS user
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT, 
	password TEXT
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

func (store *dbStore) GetMovieById(id int64) (*Movie, error) {

	var movie = &Movie{}

	err := store.db.Get(movie, "SELECT * FROM movie WHERE id=$1", id)

	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (store *dbStore) CreateMovie(m *Movie) error {

	res, err := store.db.Exec("INSERT INTO movie (title, release_date, duration, trailer_url) VALUES (?, ?, ?, ?)",
		m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)

	if err != nil {
		return err
	}

	m.ID, err = res.LastInsertId()
	return err
}

// Find User by username and password
func (store *dbStore) FindUser(username string, password string) (bool, error) {

	var count int

	err := store.db.Get(&count, "SELECT COUNT(id) FROM user WHERE username=$1 AND password=$2", username, password)

	if err != nil {
		return false, nil
	}

	return count == 1, nil
}
