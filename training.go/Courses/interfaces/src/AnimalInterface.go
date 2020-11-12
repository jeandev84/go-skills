package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	color string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	jumpHeight int
}

func (c Cat) Speak() string {
	return "Miaou"
}

// Entry point
// go run main.go
func main() {

	/*
		var a Animal
		fmt.Println(a)
	*/

	animals := []Animal{
		Dog{"white"},
		Cat{2},
	}

	for _, a := range animals {
		fmt.Println(a.Speak())
		fmt.Printf("Type of animal? %T\n", a)

		// converti a en chien (afin qu'on puisse faire des verifications)
		// t, ok := a.(Dog)
		// fmt.Printf("Type assertion value=%v, ok=%v\n", t, ok)

		if t, ok := a.(Dog); ok {
			fmt.Printf("We have a dog! color=%v\n", t.color)
		} else {
			fmt.Println("It's not a dog here...")
		}
	}

	fmt.Println("--------------------")

	for _, a := range animals {
		describeAnimal(a)
	}
}

// service interface
func describeAnimal(a Animal) {

	// "type" est un mot reserve en go pour definir le vrai type d'une varibale
	switch v := a.(type) {
	case Dog:
		fmt.Printf("We have a dog, color=%v\n", v.color)
	case Cat:
		fmt.Printf("We have a cat, jumpHeight=%v\n", v.jumpHeight)
	}
}

/*
Woof!
Type of animal? main.Dog
Type assertion value={white}, ok=true
Miaou
Type of animal? main.Cat
Type assertion value={}, ok=false
*/
