package main

import (
	"fmt"
)

type Vector struct {
	X, Y int
}

// Entry point of program
func main() {

	/* On declare une Map qui prendra en cle un entier "int" et comme valeur "boolean" */
	// var m map[int]bool = make(map[int]bool)
	// var m2 = make(map[int]bool)

	m3 := make(map[int]bool)
	fmt.Println(m3)

	vecs := make(map[Vector]string)
	fmt.Println(vecs)
}
