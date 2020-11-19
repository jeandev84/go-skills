package main

import (
	"fmt"
	"log"
	"net/http"
)

// Routes handles
func hello(w http.ResponseWriter, r *http.Request) {

	// Ecriture d'un message en reponse HTTP
	fmt.Printf("%v %v\n", r.Method, r.URL)

	// Ecrire la response retourner
	fmt.Fprintf(w, "Hello gophers from the web!")
}

func goodbye(w http.ResponseWriter, r *http.Request) {

	// Ecriture d'un message en reponse HTTP
	fmt.Printf("%v %v\n", r.Method, r.URL)

	// Ecrire la response retourner
	fmt.Fprintf(w, "Goodbye")
}

// Entry point of application
// $ go build & ./goweb
func main() {

	// Routes
	http.HandleFunc("/", hello)
	http.HandleFunc("/goodbye", goodbye)

	// Lancer le server en ecoutant au port 8000
	// go build && ./goweb
	// port :80 (site non securise http), port :443 (site securise https)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func server(port string) {

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
