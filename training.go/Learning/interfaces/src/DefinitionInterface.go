package main

import "fmt"

/*
 Une interface est un regroupement nomme  de signatures de fonctions
 Une fonction est un contract

 DEFINITION INTERFACE
 =================================================
type MyInterface interface {
	Foo() error
	Bar() string
}

var a MyInterface = ?

type MyStruct struct {

}

Implementation de methodes de MyInterface par MyStruct il suffit d'ecrire de la facon suivante:

func (m MyStruct) Foo() error {...}
func (m MyStruct) Bar() string {...}

Go deduit automatique que MyStruct est un type de MyInterface

var a MyInterface = &MyStruct{}
err := a.Foo() // appelle l'implementation de MyStruct
*/

// convention est d'ajouter le suffix "er" au nom d'interface pour deduire un comportement

// Interfaces
type Instrumenter interface {
	Play()
}

type Amplifier interface {
	Connect(amp string)
}

// Structs
type Guitar struct {
	Strings int
}

type Piano struct {
	Keys int
}

// Implementation de methodes d'interfaces
func (g Guitar) Play() {
	fmt.Printf("Tzzzzouuuuiiiing with %d strings\n", g.Strings)
}

func (g Guitar) Connect(amp string) {
	fmt.Printf("Connected to %v\n", amp)
}

// Implementation de methode Play() de l'interface Instrumenter par Piano
func (p Piano) Play() {
	fmt.Printf("Plip plip %d keys\n", p.Keys)
}

// Entry point
// go run main.go
func main() {
	var instr Instrumenter
	instr = &Guitar{5}
	instr.Play()

	instr = &Piano{88}
	instr.Play()

	// call implemented methods
	g := Guitar{12}
	g.Play()
	g.Connect("Marshall")
}
