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

func describeAnimal(a Animal) {
	switch v := a.(type) {
	case Dog:
		fmt.Printf("We have a dog, color=%v\n", v.color)
	case Cat:
		fmt.Printf("We have a cat, jumpHeight=%v\n", v.jumpHeight)
	}
}

func main() {
	var a Animal
	fmt.Println(a) // nil

	animals := []Animal{
		Dog{"white"},
		Cat{2},
	}

	for _, a := range animals {
		fmt.Println(a.Speak())
		fmt.Printf("Type of animal? %T\n", a)

		t, ok := a.(Dog)
		fmt.Printf("type assertion value=%v, ok=%v\n", t, ok)
		if ok {
			fmt.Printf("We have a dog! Color=%v\n", t.color)
		} else {
			fmt.Println("We don't have a dog here...")
		}
	}

	fmt.Println("-------------")
	for _, a := range animals {
		// a.color // impossible instruction, Animal type is not enough
		describeAnimal(a)

	}
}
