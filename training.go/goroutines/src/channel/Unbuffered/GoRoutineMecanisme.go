package main

import (
	"fmt"
	"time"
)

// Ecrie dans mon channel : (c : channel)
func ping(c chan string) {

	// on increment de 1 en 1 a chaque tour de boucle
	for i := 1; ; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
}

// Ecrire dans mon channel
func pong(c chan string) {

	// on increment de 100 a 100 a chaque tour de boucle
	for i := 100; ; i += 100 {
		c <- fmt.Sprintf("pong %v", i)
	}
}

// Lire mon channel
func print(c chan string) {

	// as While() {}
	// on lit a l'interieur de mon channel et on stock dans msg
	// Boucle infini
	for {
		msg := <-c
		fmt.Println(msg)

		// On attend une seconde entre chaque lecture
		time.Sleep(1 * time.Second)
	}
}

// GOROUTINES & CHANNELS
// Entry point of application
// $ go run main.go
func main() {

	// On declare un channel
	c := make(chan string) // est une goroutine

	// Lancer une go routine a plusieurs fonctions a la fois
	go ping(c)
	go pong(c)
	go print(c)

	// fin d'execution apres 10 seconds
	time.Sleep(10 * time.Second)
}

/*
GO ROUTINE MULTI EXECUTION
~/go/src/training.go/goroutines$ go run main.go
ping 1
pong 100
ping 2
pong 200
ping 3
pong 300
ping 4
pong 400
ping 5
pong 500
*/

/*
// $ go run main.go
func main() {

	// On declare un channel
	c := make(chan string) // est une goroutine

    // Lancer une go routine a une fonction
	go ping(c)
	go print(c)

	// fin d'execution apres 10 seconds
	time.Sleep(10 * time.Second)
}

~/go/src/training.go/goroutines$ go run main.go
ping 1
ping 2
ping 3
ping 4
ping 5
ping 6
ping 7
ping 8
ping 9
ping 10
*/
