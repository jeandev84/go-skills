package main

import (
	"fmt"
)

// Rectangle
type Rect struct {
	Width, Height int
}

// Methode de mon struct (Renvoie de la surface de notre rectangle)
func (r Rect) Area() int {
	return r.Width * r.Height
}

// Methode doublant l'aire de mon rectangle
func (r Rect) DoubleSize() {
	r.Width *= 2
	r.Height *= 2
	fmt.Println("DoubleSize()", r)
}

// Stringify like __toString() {} in php
func (r Rect) String() string {
	return fmt.Sprintf("Rect ==> width=%v, height=%v", r.Width, r.Height)
}

// Entry point of Application
func main() {

	r := Rect{2, 4}

	fmt.Printf("Rect area=%v\n", r.Area()) // ou Area(r)
	fmt.Println(r)

	r.DoubleSize()
	fmt.Println("main()", r)
}

/*
Example:

// Struct
type User struct {
    Name string
}


// Definir une Methode (Method)
// u Est un Receiver de type User
// u est un value receiver pour la methode SayHello()
// les champs de u sont accessibles pour la fonction
// Pourquoi on dit VALUE RECEIVER et non un RECEIVER ?
// c'est parce-que, le struct User est passe par copie a la fonction SayHello(),
// On dit donc c'est la valeur de User qui est passee par copie et non le User a proprement dit

La consequence est lourde car:
NB : Une methode avec une value receiver ne peut jamais modifier le struct original

func (u User) SayHello() {
   fmt.Printf("Hello %v!\n", u.Name)
}

// Tester
u := User{"Bob"}
u.SayHello() // output: Hello Bob!
*/
