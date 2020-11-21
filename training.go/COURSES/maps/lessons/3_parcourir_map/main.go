package main

import "fmt"

// Entry point of program
func main() {

	// Creer ou initializer une Map
	// m := make(map[string]int) Method 1

	m := map[string]int{
		"Bob":     10,
		"Alice":   15,
		"Bobette": 20,
		"John":    7,
	}

	/*
	 Recuperation de cle et valeur d'une MAP
	 name=John, age=7
	 name=Bob, age=10
	 name=Alice, age=15
	 name=Bobette, age=20
	*/
	for name, age := range m {
		fmt.Printf("name=%v, age=%v\n", name, age)
	}

	fmt.Print("*************************\n")

	/*
	   Recuperation de valeur de ma Map
	*/
	i := 1
	for name := range m {
		fmt.Printf("name=%v, ", name)
		m[name] = i
		i++
	}

	fmt.Println(m)
}
