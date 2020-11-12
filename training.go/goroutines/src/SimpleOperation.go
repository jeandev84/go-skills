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
	longOperation(2)

	fmt.Println("Starting second operation")
	longOperation(1)
}

/*
~/go/src/training.go/goroutines$ go run main.go
Starting first operation
Operation finished! Took 2 s
Starting second operation
Operation finished! Took 1 s
*/
