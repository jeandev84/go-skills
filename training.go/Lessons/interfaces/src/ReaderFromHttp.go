package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Entry point
// $ go run main.go
func main() {

	// Lire le contenu depuis un site web example : https://golang.org
	resp, err := http.Get("https://golang.org")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	// Creer un fichier HTML
	f, _ := os.Create("golang.html")
	defer f.Close()

	// Copy dans un fichier
	// le Body peut etre n'importe quelle chose implementant interface Reader
	io.Copy(f, resp.Body)
}

/*
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
*/
