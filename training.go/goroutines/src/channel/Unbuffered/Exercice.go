package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}
​

func main() {
	
	c := make(chan int)
	s := []int{1,25,10,3}
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}