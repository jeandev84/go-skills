package main

import "fmt"

// function no params
func printInfoNoParam() {
	fmt.Printf("Name=%s, age=%d, email=%s\n", "Bob", 10, "bob@golang.org")
}

// function with params
func printInfoParams(name string, age int, email string) {
	fmt.Printf("Name=%s, age=%d, email=%s\n", name, age, email)
}

// function average (Moyenne) x et y sont du meme types
func avg(x, y float64) float64 {
	return (x + y) / 2
}

// function avec variable de retour type
// sum est la variable retournee par la function sumNamedReturn de type int
// sum est d'ailleurs une variable qu'on peut utiliser pour des operations
func sumNamedReturn(x, y, z int) (sum int) {
	sum = x + y + z
	return sum
}

/*
Dans ce cas Go va retournee la derniere variable
qui a ete mise dans la fonction.
Methode deconseillee
func sumNamedReturn(x, y, z int) (sum int) {
	sum = x + y + z
	return
}
*/

/*
func sumNamedReturn(x, y, z int) int {
	sum := x + y + z
	return sum
}
*/

// function appelee par default quand on execute un program
func main() {
	printInfoNoParam()
	printInfoParams("Alice", 15, "alice@golang.org")

	result := avg(16.3, 25.0)
	fmt.Printf("Average result=%f\n", result)

	sum := sumNamedReturn(10, 25, 7)
	fmt.Printf("Sum result=%d\n", sum)
}
