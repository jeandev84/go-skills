package main

import (
	"fmt"
	"time"
)

func hello(c chan string) {
	c <- "Hello there!"
}

func main() {
	c := make(chan string)
	go hello(c) // The deadlock occurs only when hello is a goroutine (deadlock when *all* goroutines are asleep)
	// c <- "Hello" // Causes a deadlock because both goroutines are writing to the channel
	fmt.Println(<-c) // Fix the deadlock!
	time.Sleep(1 * time.Second)
}
