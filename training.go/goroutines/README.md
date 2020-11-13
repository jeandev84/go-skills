 # GOROUTINE 
 - Permet de faire de la programmation parallele
 Elle consiste a separer toutes les taches en petits morceaux
 en les executants part a part


``function LongOperation(i int) {``
``    // Take many seconds``
``    fmt.Println("Done", i)``
``}``

``go LongOperation(1)``
``go LongOperation(2)``
``go LongOperation(3)``
``output: Done 2, Done 3, Done 1``

GOROUTINE :
 Avec les goroutines l'ecriture et la lecture est synchrone

~/go/src/training.go/goroutines$ go run main.go
Starting first operation
Starting second operation
Operation finished! Took 1 s
Operation finished! Took 1 s

CHANNEL:

        ------ Value 1 ---- Goroutine 1
Channel ------ Value 2 ---- Goroutine 2
		------ Value 3 ---- Goroutine 3

c := make(chan <type>) // type est le type de variable qui transitera
=======================================================================

C := make(chan int)
c <- 1 // ecrire dans un channel
i := <-c // lire dans un channel, i=1

=======================================================================

func calculatePi(c chan float64) {
	// long calculation ...
	c <- resPi
}

c := make(chan float64)
go calculatePi(c) // execute dans une goroutine
pi := <-c  // bloquant juesqu'a ce que calculatePi() ecrive dans c
fmt.Println("Pi value is", pi) // output 3.1415926535...