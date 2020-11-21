package main

import (
	"fmt"

	"training.go/helloworld/data"
)

// entry point program
func main() {
	fmt.Println(data.Name, data.Age)
	/* fmt.Println(data.password) variable non accessible */
}
