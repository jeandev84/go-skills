package main

import (
	"fmt"
	"time"
)

// long operation
/*
func longOperation(done chan bool) {
	fmt.Println("Started long operation...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
	done <- true
}
*/

// Pour dire que notre operation
// pourrait etre utiliser seulement en ecriture on ecrit "chan<-"
// Si on essait de lire dans done comme ceci (<-done) cela affichera une erreure
func longOperation(done chan<- bool) {
	fmt.Println("Started long operation...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
	done <- true
}

// $ go run main.go
func main() {

	done := make(chan bool)
	go longOperation(done)
	<-done // lire la valeur de "done"
	fmt.Println("Back to main")

	/*
		Restrinct
		done := make(chan<- bool)
		go longOperation(done)
		<-done // lire la valeur de "done"
		fmt.Println("Back to main")
	*/
}

/*
sendch := make(chan<- int) // uniquement de l'envoi.
recvch := make(<-chan int) // uniquement de la reception

================================================================

func send(c chan<- int) {
	c <- 10
	i <- c // interdit
}

c:= make(chan int)
go send(c)
*/
