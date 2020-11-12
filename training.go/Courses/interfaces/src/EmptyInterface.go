package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Cooker interface {
	Cook()
}

type Chef struct {
	Person
	Stars int
}

func (c Chef) Cook() {
	fmt.Printf("I'm cooking with %v stars\n", c.Stars)
}

func proccessPerson(i interface{}) {
	switch p := i.(type) {
	case Person:
		fmt.Printf("We have a person=%v\n", p)
	case Chef:
		fmt.Printf("We have a chef=%v, Let him cook...\n", p)
		p.Cook()
	case int:
		// dosomething for type int
	default:
		fmt.Printf("Unknown type=%T, value=%v\n", p, p)
	}
}

// Entry point
// go run main.go
func main() {

	// Definir une interface Vide (Empty Interface)
	var x interface{} = Person{"Bob", 10}

	fmt.Printf("x type=%T, data=%v\n", x, x)

	// Il faut toujours indique 'ok' sinon il n'aura pas de conversion
	// Mais plutot une panique
	// Conversion d'interface en une chaine de characteres (string)
	s, ok := x.(string)
	fmt.Printf("Person as string ok? %v. value='%v'\n", ok, s)

	// Conversion d'interface en un entier
	i, ok := x.(int)
	fmt.Printf("Person as int ok? %v. value='%v'\n", ok, i)

	x = Chef{
		Stars: 4,
		Person: Person{
			Name: "Alice",
			Age:  22,
		},
	}

	proccessPerson(x)
	proccessPerson(3)
	proccessPerson("John")
}
