package main

import (
	"fmt"
	"time"
)

/*
Buffered channel work has Queue

v := make(chan int, 2) [ le channel a une capacite de 2 elements ]
v <- 10 [ non-bloquant ]
v <- 44 [ channel plein ! ]

fmt.Println(<-c)
fmt.Println(<-c)
*/

// Fonction d'ecriture dans un channel
func write(c chan string) {

	// Definir une liste de noms dans une slice
	names := []string{"Bob", "Alice", "Bobette", "John"}

	// On ecrit a chaque tour de boucle le nom courrant dans le channel
	for _, n := range names {

		c <- n
		fmt.Printf("Wrote %v to channel (len=%v)\n", n, len(c))
	}

	// Ferme le channel
	close(c)
}

// $ go run main.go
func main() {

	c := make(chan string, 3)
	fmt.Printf("Channel data: cap=%v, len=%v\n", cap(c), len(c))

	go write(c)

	// 2 seconds de traitement
	time.Sleep(2 * time.Second)

	fmt.Println("Reading ...")

	// Lire dans mon channel
	for v := range c {
		fmt.Printf("read value %v (len=%v)\n", v, len(c))
		time.Sleep(1 * time.Second)
	}
}

/*
Channel data: cap=3, len=0
Wrote Bob to channel (len=1)
Wrote Alice to channel (len=2)
Wrote Bobette to channel (len=3)
Reading ...
read value Bob (len=2)
read value Alice (len=1)
read value Bobette (len=0)
*/

// Do sleeping
/*
func sleeping(second int) {
	time.Sleep(second * time.Second)
}
*/

/*
Ce code ne compile pas:

func write(c <-chan int) {
    for i := 1; i <= 5; i++ {
        c <- i
    }
}

c := make(chan int)
go write(c)

time.Sleep(2 * time.Second)
fmt.Println(<- c)
fmt.Println(<- c)
fmt.Println(<- c)
*/
