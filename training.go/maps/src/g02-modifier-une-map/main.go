package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	// assign value to key
	m["hello"] = 5
	m["goodbye"] = 10
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	// retrieve value
	fmt.Printf("key=hello, value=%v\n", m["hello"])
	i := m["goodbye"]
	fmt.Printf("i=%v\n", i)

	// test presence
	j, ok := m["helo"]
	fmt.Printf("j=%v, ok=%v\n", j, ok)

	// test only the key
	m["beatles"] = 2
	if _, ok := m["beatles"]; ok {
		fmt.Println("beatles key exist!")
	}

	// delete key
	// does nothing if no key
	delete(m, "beatles")

	// reference type !
	m2 := m
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))
	m2["help"] = 44
	fmt.Printf("m content %v (len=%v)\n", m, len(m))
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))
}
