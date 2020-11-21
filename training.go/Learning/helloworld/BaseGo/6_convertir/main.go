package main

import (
	"fmt"
)

func main() {

	// Declaration de variable
	var percentage float64 = 2.5

	// Augmentation de pourcentage
	// percentage = percentage + 34 or percentage += 34
	percentage += 34

	// Current percentage 2.0%
	fmt.Printf("Current percentage %f%%\n", percentage)
	fmt.Printf("Int value=%d%%\n", int(percentage))

	n := 42
	f := float64(n) + 0.42
	fmt.Printf("float=%f\n", f)
}
