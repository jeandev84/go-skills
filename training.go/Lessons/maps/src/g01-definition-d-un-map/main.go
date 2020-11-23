package main

import (
	"fmt"
)

type Vector struct {
	X, Y int
}

func main() {
	// long
	// var m map[int]bool = make(map[int]bool)

	// short
	// var m = make(map[int]bool)

	// shortest
	m := make(map[int]bool)
	fmt.Println(m)

	// key = struct
	vecs := make(map[Vector]string)
	fmt.Println(vecs)
}
