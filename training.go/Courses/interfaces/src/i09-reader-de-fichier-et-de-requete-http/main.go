package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// read to string
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// read to file
	f, _ := os.Create("golang.html")
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
}
