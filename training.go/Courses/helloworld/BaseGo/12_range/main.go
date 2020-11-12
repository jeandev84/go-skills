package main

import (
	"fmt"
)

/*
Range signifit interval
Il permet d'iterer sur les elements d'une structure donnees
typiquement des tableaux et des slices

for <index>, <value> := range <data> {
   // code
}

nums := []int{15, -2}
for i, num := range nums {
	fmt.Printf("nums[%d]=%d\n", i, num)
}

// output: nums[0] = 15, nums[0] = 15, nums[1] = -2
*/

func main() {

	names := []string{"Bob", "Alice", "Bobette", "John"}

	for i, n := range names {
		fmt.Printf("Username=%s index=%d)\n", n, i)
	}

	// range en ignorant l'index en ecrivant underscore "_"

	for _, c := range "golang" {
		/* fmt.Printf("%v\n", c) Affiche la valeur de c en memoire */
		fmt.Printf("%v\n", string(c))
	}
}
