package main

import (
	"fmt"
)

/*
Pointeur
- C'est un mecanisme royal pour partager et modifier la memoire
- Un pointeur est une variable qui stock l'addresse d'une autre variable
- Les Pointeurs En Go, nous permettent de modifier les parametres des fonctions
- C'est exactement ce qu'on aimerait faire avec la methodes

============================
Memoire
   x     s    p     i
  -42    Bob  @1   -42
   @1    @2   @3    @4
=============================

x := -42
s := "Bob"
p := &x    // Creation d'un pointeur vers la variable x

(On recupere le contenu de ce que porte p, c'est a dire le contenu de @1)
i := *p    // Dereferencement de p pour recuperer la valeur de x

*/

// Update value
func UpdateValue(val string) {
	val = "value"
}

// Update pointer (By Deferencement type)
func UpdatePointer(ptr *string) {
	*ptr = "pointer" // Dereference la valeur dont il pointe (decode)
}

// Entry point of Application
func main() {

	// Declaration de pointeur p. (Cas int)
	// pointeur sur un integer
	// p va contenir l'address de la variable i
	i := 1
	var p *int = &i

	fmt.Printf("i=%v\n", i)
	fmt.Printf("p=%v\n", p)   // reference p (sauvegarder l'addresse d'une valeur en memoire)
	fmt.Printf("*p=%v\n", *p) // deference p (recuperer la valeur de l'addresse sauvegardee)
	fmt.Println("---------------")

	// Declaration de pointeur s. (Cas string)
	// var sPtr *string = &s  (Longue syntaxe)
	s := "Bob"
	sPtr := &s  // sauvegarde l'addresse en memoire
	s2 := *sPtr // recupere la valeur de l'addresse en memoire var s2 string = *sPtr

	fmt.Println("String pointer")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("---------------")

	// Mise ajour des referencements
	// va modifier le contenu de l'addresse qui est contenu dans mon pointeur
	*sPtr = "Alice"

	fmt.Println("Derefence and update")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("---------------")

	UpdateValue(s)
	fmt.Println("Func UpdateVal()")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Println("---------------")

	UpdatePointer(&s) // or UpdatePointer(sPtr)
	fmt.Println("Func UpdatePointer()")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Println("---------------")
}
