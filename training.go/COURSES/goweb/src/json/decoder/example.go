package main

import (
	"encoding/json"
	"log"
)

// User struct
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Entry point
func main() {
	b := []byte(`{"name":"Bob","password:"secret", "email":"bob@golang.org"}`)

	var u User

	// Decoder du JSON
	err := json.Unmarshal(b, &u)

	if err != nil {
		log.Fatal(err)
	}
}
