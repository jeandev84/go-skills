package main

import (
	"encoding/json"
	"fmt"
)

// User entity
// utiliser les structs tags pour formatter du json `json:"nom_en_minuscule"`
// Back tild ``
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {

	u := User{
		Name:     "Bob",
		Password: "secret",
		Email:    "bob@golang.org",
	}

	b, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	// {"Name": "Bob", "Password": "secret", "Email": "bob@golang.org"}
	fmt.Println(string(b))
}
