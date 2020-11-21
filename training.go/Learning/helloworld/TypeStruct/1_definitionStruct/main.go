package main

/*
 DECLARATION DE STRUCT
*/
// Struct Address de mes utilisateurs
type Address struct {
	street, city string
}

// Struct Person
type Person struct {
	Name string
	Age  int
	Addr Address
}

// Entry point of application
func main() {

	// Declaration de la variable
	var p Person

	p.Name = "Bob"
	p.Addr.city = "Lyon"

}
