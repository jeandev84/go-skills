package main

import (
	"fmt"
)

// Entry point of program
func main() {

	// Creer ma map
	m := make(map[string]int)

	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	// Assigner une valeur a ma cle dans ma Map
	m["hello"] = 5
	m["goodbye"] = 10

	// Map content map[goodbye:10 hello:5] (len=2)
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	// Recuperer une valeur de cle dans une Map
	// key=hello, value=5
	fmt.Printf("key=hello, value=%v\n", m["hello"])

	// Assigner la valeur a une cle de Map
	i := m["goodbye"]
	fmt.Println(i)

	// Tester la presence d'une key dans le Map (ok est oblige)
	// On verifie si la cle "foo" existe dans ma Map
	j, ok := m["foo"]
	fmt.Printf("j=%v, ok=%v\n", j, ok)

	// Tester si la cle exist e afficher un message
	m["beatles"] = 2

	// ;ok --> ;ok == true
	if _, ok := m["beatles"]; ok {
		fmt.Println("beatles key exist")
	}

	// Supprimer une cle dans une Map
	delete(m, "beatles")
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	// Assignation
	m2 := m
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))

	m2["help"] = 44
	fmt.Printf("m content %v (len=%v)\n", m, len(m))
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))
}

/*
Map content map[] (len=0)
Map content map[goodbye:10 hello:5] (len=2)
key=hello, value=5
10
j=0, ok=false
beatles key exist
Map content map[goodbye:10 hello:5] (len=2)
m2 content map[goodbye:10 hello:5] (len=2)
m content map[goodbye:10 hello:5 help:44] (len=3)
m2 content map[goodbye:10 hello:5 help:44] (len=3)
*/
