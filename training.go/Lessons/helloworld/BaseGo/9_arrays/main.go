package main

import "fmt"

/*
var nom[taille]type

Example:
- Declaration un tableau de nom "t" de taille 5 de nombres entiers
var t[5]int

- Affectation (Set)
t[3] = 12

- output: 12 ===> Lecture (Get)
fmt.Println("", t[3])

- Recuperer la taille d'un tableau
*/

func main() {

	// liste de noms
	var names [3]string

	// %v (value) permet d'afficher le type qui correspond
	// on peut mettre un string, int, tableau, booleen Go retrouvera le type de valeur
	fmt.Printf("names=%v (len=%v)\n", names, len(names))

	// Affectation de valeurs a mon tableau
	names[0] = "Bob"
	names[2] = "Alice"
	fmt.Printf("names=%v (len=%v)\n", names, len(names))
	fmt.Printf("names[2]=%v\n", names[2])

	// Declaration de nombres impaires:
	odds := [4]int{1, 3, 5, 7} // odds := [4]int{1, 3}
	fmt.Printf("odds=%v (len=%d)\n", odds, len(odds))
}
