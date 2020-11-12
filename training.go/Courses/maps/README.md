# DEFINITION MAP
Une Map permet d'associer a une cle une valeur comme les arrays
Syntaxe
``var m map[KeyType]ValueType``
``var m map[TypeDeLaCle]TypeDeLaValeur``

``Example de declaration de Maps: var m map[string]int``
``m = make(map[string]int)``


# MANIPULER UNE MAP
- Assigner une valeur
- Recuperer une valeur
- Tester la presence
- Supprimer couple cle / valeur


# PARCOURIR UNE MAP
package main

import "fmt"

// Entry point of program
func main() {

	// Creer ou initializer une Map
	// m := make(map[string]int) Method 1

	m := map[string]int{
		"Bob":     10,
		"Alice":   15,
		"Bobette": 20,
		"John":    7,
	}

	/*
	 Recuperation de cle et valeur d'une MAP
	 name=John, age=7
	 name=Bob, age=10
	 name=Alice, age=15
	 name=Bobette, age=20
	*/
	for name, age := range m {
		fmt.Printf("name=%v, age=%v\n", name, age)
	}

	fmt.Print("*************************\n")

	/*
	   Recuperation de valeur de ma Map
	*/
	i := 1
	for name := range m {
		fmt.Printf("name=%v, ", name)
		m[name] = i
		i++
	}

	fmt.Println(m)
}


# STRUCT ET MAP
Pour être utilisé en tant que clé, une valeur doit être "comparable" (opérateur ==). En Go, cela inclus tout sauf : les map, slice, fonction.

- package main

import (
	"fmt"
)

// Struct
type User struct {
	name string
}

type Key struct {
	ID   int
	Name string
}

// Entry point of program
func main() {

	// CAS Struct
	m := map[string]*User{
		"HR":  {"Bob"},
		"CEO": {"Alice"},
	}

	fmt.Println(m["HR"])

	hr := m["HR"]
	hr.name = "John"

	fmt.Println(m["HR"])

	// CAS MAP
	m["CFO"] = &User{"Bobette"}

	res := make(map[Key]string)
	res[Key{1, "aze"}] = "file1"

	k := Key{2, "ert"}
	res[k] = "file2"

	// Delete
	delete(res, Key{1, "aze"})
	fmt.Println(res)
}

/*
&{Bob}
&{John}
map[{2 ert}:file2]
*/

# Example
m := map[string]bool{
    "Beatles": true,
    "Rolling Stones": true,
    "Led Zeppelin": false,
}
for band, present := range m {
    if band == "Rolling Stones" {
        delete(m, band)
    }
    fmt.Printf("%v=%v\n", band, present)
}