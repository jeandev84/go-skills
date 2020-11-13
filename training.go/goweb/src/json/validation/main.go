package main

import (
	"encoding/json"
	"fmt"
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

// PasswordJsonBody
// Si le champ est ecrit en minuscule 'json' ne les executera pas car il les considera
// comme des champs caches
// UserIndex : l' index de l'utilisateur pour lequel on veut changer le mot de passe
// OldPassword: l'ancien mot de passe
// NewPassword: le nouveau mot de passe
// NewRepeatPassword: le nouveau mot de passe repete
// Nouveau mot de passe
/*
{
	 "user_index": 0,
	 "old_password": "secret",
	 "new_password": "testsecret",
	 "new_password_repeat": "testsecretERROR"
}
*/
type PasswordJsonBody struct {
	UserIndex         int    `json:"user_index"`
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	NewRepeatPassword string `json:"new_password_repeat"`
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

// Update user password (change password)
func updatePassword(w http.ResponseWriter, r *http.Request) {

	var p PasswordJsonBody

	// Decode le body de la request (get request content)
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {

		// Repondre au navigateur qu'il y a une erreure (requette malformee)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Fprintf(w, "Parsed struct %v", p)
	log.Printf("Parsed struct %v", p)

	// verifions si l'index de l'utilisateur est un index valid
	if p.UserIndex < 0 || p.UserIndex > len(users)-1 {
		msg := fmt.Sprintf("Invalid index. got user_index=%v. valid range=[0,%v]", p.UserIndex, len(users)-1)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Si l'index d'utilisateur est valide cela veut dire que on va chercher l'utilsateur associe
	u := users[p.UserIndex]

	// Verifions si les mots de passe correspondent
	if u.Password != p.OldPassword {
		http.Error(w, "Old password do not match", http.StatusBadRequest)
		return
	}

	// Verifions si les deux passwords correspondent
	if p.NewPassword != p.NewRepeatPassword {
		http.Error(w, "New passwords do not match", http.StatusBadRequest)
		return
	}

	// Si tout va bien on peut mettre ajour le mot de passe
	u.Password = p.NewPassword

	// Envoyer le message que le password a ete modifier
	fmt.Fprintf(w, "Pass updated")

}

// List users handles
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
	http.HandleFunc("/update_password", updatePassword)

	// Server web
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
