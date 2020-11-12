package main

import (
	"fmt"
)

// User struct
type User struct {
	Name  string
	Email string
	//IsAdmin bool
}

// Embedded struct (Heritage)  Un Admin est un utilisateur User
// Un admin a un niveau de securite Level
type Admin struct {
	User
	Level int
}

// Entry point of Application
func main() {
	u := User{
		Name:  "Bob",
		Email: "bob@golang.org",
	}

	fmt.Printf("User=%v\n", u)

	/*
		a := Admin{
			Level: 2,
		}

		a.Name = "Alice"
		a.Email = "alice@golang.org"

		fmt.Printf("Admin=%v\n", a)
	*/

	a := Admin{
		Level: 2,
		User: User{
			Name:  "Alice",
			Email: "alice@golang.org",
		},
		/*
			Interdit d'ecrire directement Name et Email dans Admin car il faut obeir a la structure
			Name:  "Alice",
			Email: "alice@golang.org",
		*/
	}

	fmt.Printf("Admin=%v\n", a)

	// On a access au proprietes de User
	fmt.Printf("Admin name=%v, level=%v\n", a.Name, a.Level)
}

/*
# Embedded Struct

``Identifie le lien entre 2 Structs``


Quel est le lien qui est exprime entre les 2 structs

type Address struct {
     City string
}


type User struct {
    Name string
    Addr Address
}


Relation de possession Owner Ship
- un User a une Address (OneToOne)

Un XXX est un YYY (Heritage)


Go prefere la composition avec l'Embedded Struct
Embedded Struct signifit on va inclure un Struct a l'interieur d'un autre
Ici Address est inclus dans User et on peut l'utiliser
Address est embarque dans le type User (On peut dire User est une Address)

type Address struct {
     City string
}


type User struct {
    Name string
    Address // pas nom de variables
}

var u User
u.City = "Londres" // City est directement accessible
*/
