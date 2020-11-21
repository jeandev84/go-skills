package main

import (
	"fmt"
)

type User struct {
	name string
}

type Key struct {
	ID   int
	Name string
}

func main() {
	// Using a pointer or not has a big impact!
	m := map[string]*User{
		"HR":  {"Bob"},
		"CEO": {"Alice"},
	}

	fmt.Println(m["HR"])

	hr := m["HR"]
	hr.name = "John"
	fmt.Println(m["HR"])

	// when using a pointer = just copy the pointer
	// when using a value = copy the entire struct
	m["CFO"] = &User{"Bobette"}

	res := make(map[Key]string)
	res[Key{1, "aze"}] = "file1"

	k := Key{2, "ert"}
	res[k] = "file2"
	fmt.Println(res)
	delete(res, Key{1, "aze"})
	fmt.Println(res)
}
