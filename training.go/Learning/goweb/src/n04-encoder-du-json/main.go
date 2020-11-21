package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country,omitempty"`
}

type User struct {
	Name     string  `json:"name"`
	Password string  `json:"-"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
}

func users(w http.ResponseWriter, r *http.Request) {
	users := []User{
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
				Street:  "42 rue elle",
				City:    "Paris",
				Country: "",
			},
		},
	}

	b, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	// Associe un chemin à une fonction
	http.HandleFunc("/users", users)

	// Création du serveur HTTP écoutant sur le port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
