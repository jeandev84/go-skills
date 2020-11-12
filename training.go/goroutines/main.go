package main

import (
	"fmt"
	"time"
)

// Longue operation
func longOperation(duration int) { // duration in seconds

	// Endormir le process
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("Operation finished! Took %d s\n", duration)
}

// GOROUTINES
// Entry point of application
// $ go run main.go
func main() {

	fmt.Println("Starting first operation")
	go longOperation(1)

	fmt.Println("Starting second operation")
	go longOperation(2)

	time.Sleep(2 * time.Second)
}

/*
~/go/src/training.go/goroutines$ go run main.go
Starting first operation
Starting second operation
Operation finished! Took 1 s
Operation finished! Took 1 s
*/
