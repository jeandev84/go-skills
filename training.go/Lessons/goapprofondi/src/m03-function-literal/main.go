package main

import (
    "fmt"
    "sort"
)

func main() {
    // Définition simple
    func() {
		fmt.Println("Hello!")
	}()

    // Support d'un argument
    func(msg string) {
		fmt.Println(msg)
	}("Hello with arg")

    // Utilisation pour trier une liste
    people := []string{"Alice", "Bob", "Dave"}
    sort.Slice(people, func(i, j int) bool {
        return len(people[i]) < len(people[j])
    })
    fmt.Println(people)

    // Le lambda peut être stocké dans une variable
    // La fonction utilise une closure
    // Elle accède a la variable people définie en dehors de son scope
    less := func(i, j int) bool {
        return len(people[i]) < len(people[j])
    }
    sort.Slice(people, less)
}

