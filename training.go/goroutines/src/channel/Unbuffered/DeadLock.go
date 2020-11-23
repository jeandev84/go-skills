package main

import (
	"fmt"
	"time"
)

// Fonction hello()
func hello(c chan string) {

	// la notation <- permet d'ecrire dans notre channel "c"
	c <- "Hello there!"

	fmt.Println("hello() finished")

	// Lecture de mon channel
	fmt.Println(<-c)
}

// $ go run main.go or $ go build && ./goroutines or $ ./goroutines
// ecrire  dans un channel : c := make(chan string); c <- "Message from main"
// lecture dans un channel : fmt.Println(<-c)
func main() {

	c := make(chan string)
	go hello(c)

	// Lecture dans un channel (Permet de resoudre le probleme de "deadlock")
	fmt.Println(<-c)

	// ecrire dans un channel
	c <- "Message from main"

	time.Sleep(1 * time.Second)
}

/*
sendch := make(chan<- int) // uniquement de l'envoi.
recvch := make(<-chan int) // uniquement de la reception
*/
