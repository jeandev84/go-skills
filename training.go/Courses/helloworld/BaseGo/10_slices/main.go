package main

import "fmt"

/*
un slice va pouvoir augmenter en fonction du contenu

SLICE = Tranche
Un Slice represente une tranche d'un tableau

Un Slice est une "vue" sur le tableau qui stock les valeurs
Lorsqu'on modifie un slice on modifit le tableau qui est derriere


Syntaxe:

s := make([]type, taille, capacite)
. Taille:
  correspond au nombre d'elements du slice

. Capacite[facultatif]:
  nombre d'elements du tableau par default il prendra la taille du tableau
  Losqu'on depasse la taille du tableau la Capacite se double


  Example:
  Initialisation de slice de taille 3
  s := make([]int, 3)
  s[0] = -3
  len(s) // 3 function len() permet de determiner la taille du slice
  cap(s) // 3 function cap() permet de determiner la capicite du slice

  Ajout de valeur dans un slice
  s := make([]int, 3)
  s = append(s, 12)


  Documentation sur les slices
  https://blog.golang.org/go-slices-usage-and-internals
*/

func main() {

	/*
	 []int (type de valeurs entiers)
	  2 taille du tableau
	  3 capacite peut porter le tableau
	*/
	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = -1
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	nums = append(nums, 64)
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	nums = append(nums, -8)
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	/*
	 Methode de collection d'elements
	*/
	letters := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Println("---------------------------------------")

	letters = append(letters, "!")
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))

	// Recuperation des 3 premieres lettres de l'indice 0 a 2
	sub1 := letters[:2] // letters[0:2] (debut a un indice precis)
	sub2 := letters[2:] // letters[2:len(letters)] de 2 jusqu'a la fin
	fmt.Printf("%v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Println("---------------------------------------")

	// Modifier un slice
	sub2[0] = "UP"
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Println("---------------------------------------")

	// Copier un slice
	subCopy := make([]string, len(sub1))
	copy(subCopy, sub1)

	subCopy[0] = "DOWN"
	fmt.Printf("COPY %v (len=%d, cap=%d)\n", subCopy, len(subCopy), cap(subCopy))
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Println("---------------------------------------")
}
