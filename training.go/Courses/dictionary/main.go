package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/dictionary/dictionary"
)

/*
Install package:
On indique (3) petits points (...) pour telecharger les 3 sous petites dependences du package
$ go get github.com/dgraph-io/badger/...
Entry point
Le package "flag" permet de travailler avec la CLI en go
https://golang.org/flag/
*/
func main() {

	// CLI (Command Line Interface)
	// number := flag.String("number", "list", "Action to perform on the dictionary")
	action := flag.String("action", "list", "Action to perform on the dictionary")

	// New Dictionary
	d, err := dictionary.New("./badger")

	handleErr(err)

	defer d.Close()

	flag.Parse()

	// Testing action
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknown action: %v\n", *action)
	}

}

// Action list dictionary
func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("Dictionary content")

	for _, word := range words {
		fmt.Println(entries[word])
	}
}

// Action add dictionary
func actionAdd(d *dictionary.Dictionary, args []string) {

	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)

	handleErr(err)
	fmt.Printf("'%v' added to the dictionary\n", word)
}

// Action define dictionary
func actionDefine(d *dictionary.Dictionary, args []string) {

	word := args[0]
	entry, err := d.Get(word)

	handleErr(err)
	fmt.Println(entry)
}

// Action define dictionary
func actionRemove(d *dictionary.Dictionary, args []string) {

	word := args[0]
	err := d.Remove(word)

	handleErr(err)
	fmt.Printf("'%v' was removed from the dictionary\n", word)
}

// private method handleErr() for displaying error
// given parameter "err"
func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}
}
