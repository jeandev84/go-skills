package main

import (
	"fmt"
	"sort"
)

func main() {

	// Definition simple
	func() {
		fmt.Println("Hello!")
	}()

	// Support d'un argument
	func(msg string) {
		fmt.Println(msg)
	}("Hello with args")

	// Utilisation pour trier une liste
	people := []string{"Alice", "Bob", "Dave"}
	// sort.Slice(people, func(i, j int) bool {
	// 	return len(people[i]) < len(people[j])
	// })

	// fmt.Println(people)

	// Le lambda peut etre stocke dans une variable
	/*
		var less = func(i, j int) bool {
			return len(people[i]) < len(people[j])
		}
	*/

	// La fonction utilise une closure
	// Elle accede a la variable people definie en dehors de son scope
	less := func(i, j int) bool {
		return len(people[i]) < len(people[j])
	}
	sort.Slice(people, less)
	fmt.Println(people)
}
