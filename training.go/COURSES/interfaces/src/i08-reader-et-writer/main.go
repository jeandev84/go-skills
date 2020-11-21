package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// r := strings.NewReader("Hello there Gophers! Readers are awesome")
	r, err := os.Open("test.txt")
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Printf("Error in Reader: %v\n", err)
		return
	}

	fmt.Println(string(buf))
}
