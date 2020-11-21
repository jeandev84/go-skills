package main

import (
	"fmt"
	"log"
	"net/http"
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

	/*
		code de test
		movies, err := srv.store.GetMovies()

		if err != nil {
			return err
		}

		fmt.Printf("movies=%v\n", movies)
	*/

	defer srv.store.Close()

	// On delegue l'ensemble des routes a mon routeur
	http.HandleFunc("/", srv.serveHTTP)

	log.Printf("Serving HTTP on port 9000")

	err = http.ListenAndServe(":9000", nil)

	if err != nil {
		return err
	}

	return nil
}
