package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
			Street:  "42 rue elle",
			City:    "Paris",
			Country: "",
		},
	},
}

type PasswordJsonBody struct {
	UserIndex         int    `json:"user_index"`
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	NewPasswordRepeat string `json:"new_password_repeat"`
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
	var p PasswordJsonBody

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Update password parsed: %v\n", p)

	if p.UserIndex < 0 || p.UserIndex > len(users)-1 {
		msg := fmt.Sprintf("Invalid index. got user_index=%v, valid range=[0, %v]", p.UserIndex, len(users)-1)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	u := users[p.UserIndex]

	if u.Password != p.OldPassword {
		http.Error(w, "Old password do not match", http.StatusBadRequest)
		return
	}

	if p.NewPassword != p.NewPasswordRepeat {
		http.Error(w, "New passwords do not match", http.StatusBadRequest)
		return
	}

    u.Password = p.NewPassword

	fmt.Fprintf(w, "Password updated")
}

func main() {
	// Associe un chemin à une fonction
	http.HandleFunc("/update_password", updatePassword)

	// Création du serveur HTTP écoutant sur le port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
