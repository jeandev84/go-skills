package main

import (
	"fmt"
	"time"
)

/*
Select est comme un switch
qui permet de travailler avec plusieurs channel a la fois

c1 := make(chan int)
c2 := make(chan int)


select {
case v := <- c1:
   fmt.Println("Received value", v)

case <- c2:
	fmt.Println("Something happened en c2!")
}
*/

// Message from client1
func client1(c chan string) {

	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("Message from client1 => %v", i)
		time.Sleep(1500 * time.Millisecond)
	}
}

// Message from client2
func client2(c chan string) {

	for i := 10; i < 15; i++ {
		c <- fmt.Sprintf("Message from client2 => %v", i)
		time.Sleep(3000 * time.Millisecond)
	}
}

// System socket
// $ go run main.go
func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go client1(c1)
	go client2(c2)

	maxEmpty := 10
	currEmpty := 0

	for currEmpty <= maxEmpty {

		// On attend une seconde pour afficher
		time.Sleep(1000 * time.Millisecond)

		select {
		case v := <-c1:
			fmt.Println("Received from client1: ", v)
		case v := <-c2:
			fmt.Println("Received from client2: ", v)
		default:
			fmt.Println("No value received")
			currEmpty++
		}
	}
}

/*
Received from client1:  Message from client1 => 0
Received from client1:  Message from client1 => 1
Received from client1:  Message from client1 => 2
Received from client1:  Message from client1 => 3
Received from client1:  Message from client1 => 4
*/

/*
Received from client2:  Message from client2 => 10
Received from client1:  Message from client1 => 0
No value received
Received from client2:  Message from client2 => 11
Received from client1:  Message from client1 => 1
No value received
Received from client2:  Message from client2 => 12
Received from client1:  Message from client1 => 2
No value received
Received from client1:  Message from client1 => 3
Received from client2:  Message from client2 => 13
Received from client1:  Message from client1 => 4
No value received
Received from client2:  Message from client2 => 14
No value received
No value received
No value received
No value received
No value received
No value received
No value received
*/
