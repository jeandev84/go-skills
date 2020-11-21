package main

/*
fmt = format
*/

import (
	"fmt"
	"strings"

	"training.go/helloworld/input"
)

func main() {

	// input mouse (role)cd
	input.Mouse()
	// strings.ToUpper () mettre un mot en majuscule
	fmt.Println(strings.ToUpper("Hello Gophers! This is a message from the Golang Course"))
}
