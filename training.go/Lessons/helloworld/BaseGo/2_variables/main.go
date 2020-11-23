// name of package
package main

// import packages
import (
	"fmt"
)

// declare variables outside functions
// generation on declare les variables sans initialisation en dehors des fonctions
var html = "<h1>"
var pi float32 = 3.14

// entry point program
func main() {

	// var radius = 2.4

	// Declaration de variable simple
	var age int = 20
	fmt.Println(age)

	// Declaration de plusieurs variables en une ligne
	var weight, height int = 80, 190
	fmt.Println(weight, height)

	// Declaration simplifie de variable
	var name = "Bob"
	fmt.Println(name)

	// Declaration de variable compact
	// On declare et on affecte une valeur
	email := "bob@go.org"
	fmt.Println(email)

	/*
		email := "alice@go.org"
		fmt.Println(email)
	*/

	pi := 3.14
	fmt.Println(pi)
}
