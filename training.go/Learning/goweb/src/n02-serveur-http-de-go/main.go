package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// Ecriture d'un message en réponse HTTP
	fmt.Printf("%v %v\n", r.Method, r.URL)
	fmt.Fprintf(w, "Hello gophers!")
}

func main() {
	// Associe un chemin à une fonction
	http.HandleFunc("/", hello)

	// Création du serveur HTTP écoutant sur le port 8000
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
