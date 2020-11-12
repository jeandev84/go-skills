package main

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
