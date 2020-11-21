package main

import (
	"fmt"
	"os"
)

// Entry point of server
func main() {

	fmt.Println("GoFlix")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Lancer l' application
	// $ go build && ./goflix
}

// Fonction d'execution de mon serveur
func run() error {

	srv := newServer()
	srv.store = &dbStore{}
	err := srv.store.Open()

	if err != nil {
		return err
	}

	movies, err := srv.store.GetMovies()

	if err != nil {
		return err
	}

	fmt.Printf("movies=%v\n", movies)

	defer srv.store.Close()
	return nil
}
