package main

import (
	"fmt"
	"time"
)

// Read name for example
func reader(c chan string) {

	fmt.Println("Start read")
	defer fmt.Println("Stop read")

	for n := range c {
		fmt.Println(n)
	}
}

// $ go run main.go
func main() {

	// Create channel
	c := make(chan string)

	// Go routine channel
	go reader(c)

	// Write in the channel
	c <- "Bob"
	c <- "Alice"

	// Close channel
	close(c)

	/* c <- "Bobette" ! interdit d'ecire dans un channel ferme */

	time.Sleep(5 * time.Second)
}

/*
c := make(chan int)
c <- 1
c <- 2
c <- 3

for v := range c {
	fmt.Println(v) // output : 1, 2, 3
}
*/
