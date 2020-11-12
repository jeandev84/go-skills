package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Entry point
// go run main.go
func main() {

	r := strings.NewReader("Hello Gophers! Readers are awesome")
	buf, err := ioutil.ReadAll(r)

	if err != nil {
		fmt.Printf("Error in Reader: %v\n", err)
		return
	}

	fmt.Println(string(buf))
}

/*
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
*/
