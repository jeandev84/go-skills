package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hello handle
func hello(w http.ResponseWriter, r *http.Request) {

	// Ecriture d'un message en reponse HTTP
	fmt.Printf("%v %v\n", r.Method, r.URL)

	// Ecrire la response retourner
	fmt.Fprintf(w, "Hello gophers from the web!")
}

// Goodbye handle
func goodbye(w http.ResponseWriter, r *http.Request) {

	// Ecriture d'un message en reponse HTTP
	fmt.Printf("%v %v\n", r.Method, r.URL)

	// Ecrire la response retourner
	fmt.Fprintf(w, "Goodbye")
}

// Search handle
// /search?term=something
// une requette ressemble a cela : t: term, p: page
// http://localhost:8000/search?t=go&p=1
func search(w http.ResponseWriter, r *http.Request) {

	// Recuperation des parametres venant de l' URL
	t := r.URL.Query().Get("t")
	p := r.URL.Query().Get("p")

	fmt.Fprintf(w, "Searching for term=%v. Current page=%v", t, p)

}

// Login handle
func login(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/login.html")
		return
	case "POST":
		// get les parametres parses dans le formulaire a l'aide de r.ParseForm()
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() failed err=%v", err)
			return
		}

		fmt.Fprintf(w, "Go login POST. value=%v\n", r.PostForm)

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "Go" && password == "rocks" {

			fmt.Fprintf(w, "You are now logged\n")

		} else {

			fmt.Fprintf(w, "Wrong username / password\n")
		}
	}
}

// Entry point of application
// JSON (Javascript Object Notation)
// $ go build & ./goweb
func main() {

	// Associe un chemin a une fonction
	http.HandleFunc("/", hello)
	http.HandleFunc("/search", search)
	http.HandleFunc("/login", login)

	// Creation du serveur HTTP ecoutant sur le port 8000
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

	/* server("8000") */
}

// Server
func server(port string) {

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
