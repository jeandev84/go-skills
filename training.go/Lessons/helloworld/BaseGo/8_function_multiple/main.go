package main

import (
	"fmt"
	"strings"
)

/*
Function public ramenant une chaine de characteres en minuscule
et renvoyer la longueur de cette chaine de characteres
*/
func ToLowerStr(name string) (string, int) {
	return strings.ToLower(name), len(name)
}

/*
Entry point
*/
func main() {

	// Recuperation des 2 valeurs renvoyee par la fonction ToLowerStr()
	lowerName, len := ToLowerStr("ALICE")
	fmt.Printf("%s (len=%d)\n", lowerName, len)

	// _ (on ecrit underscore pour ignorer la premiere valeure qui a ete retournee)
	// Ici on met egale car la variable "len" a ete declaree
	_, len = ToLowerStr("Bob")
	fmt.Printf("%s (len=%d)\n", "bob", len)

	// Declaration d'une nouvelle variable
	bobLower, len := ToLowerStr("Bob")
	fmt.Printf("%s (len=%d)\n", bobLower, len)

}

/*
func Max(a, b int) (int, int) {
    if a > b {
        return a, b
    } else {
        return b, a
    }
}
​
num1, num2 := Max(6, 42)
fmt.Printf("num1=%d, num2=%d\n", num1, num2)
*/
