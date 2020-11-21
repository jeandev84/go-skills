package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Address
// Le mot cle "omitempty" : siginifit si le champ n'est pas renseigne on ne l'affiche pas
// Le package json va omettre le champ "country" si jamais Country n'est pas remplit
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country,omitempty"`
}

// User (utilisateur)
// json:"-" ignorer parametre
type User struct {
	Name     string  `json:"name"`
	Password string  `json:"-"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
}

// variable globale users
var users = []User{
	{
		Name:     "Bob",
		Password: "secret",
		Email:    "bob@golang.org",
		Address: Address{
			Street:  "15 rue Hade",
			City:    "Paris",
			Country: "France",
		},
	},
	{
		Name:     "Alice",
		Password: "supersecret",
		Email:    "alice@golang.org",
		Address: Address{
			Street:  "42 rue Elle",
			City:    "Paris",
			Country: "", // On omet le pays
		},
	},
}

// Handles
func usersList(w http.ResponseWriter, r *http.Request) {

	// Encodage du tableau de users (User en json)
	// data: users, prefix: "", indemptation: "  " (2 espaces)
	// b: bytes

	// b, err := json.Marshal(users) Ici pas d'indemptation!!! le json sera mal formatte
	b, err := json.MarshalIndent(users, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	// Headers format json (en-tete)
	w.Header().Set("Content-Type", "application/json")

	// Envoie de donner
	w.Write(b)
}

/*
 Response
*/
func respond(w http.ResponseWriter, headers map[string]string) {

	for key, value := range headers {
		w.Header().Set(key, value)
	}

}

// Entry point of application
// JSON (Javascript Object Notation)
// $ go build & ./goweb
func main() {

	// Routes
	http.HandleFunc("/users", usersList)

	// Server web
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
