package main

import "fmt"

/*
"defer" permet d'executer un code au moment ou on sort de la fonction

func main() {
	f := os.OpenFile("foo.txt")

	if condition1 {
		return // Oops...! pas de close ici!
	}
	// code
	f.Close()
}


func main() {

	J'ouvre mon fichier
	f := os.OpenFile("foo.txt")

	S'executera quand on sort de main()
	defer f.Close()

	if condition1 {
		return // Oops...! pas de close ici!
	}

}
*/

func start() {
	fmt.Println("Start")
}

func finish() {
	fmt.Println("Finish")
}

// Entry point of program
func main() {

	start()

	// defer - afin que la fonction soit appelee a la fin d'execution
	// LIFO = Last In First Out
	defer finish()

	// finish() .5
	// Goodbye Bob .4
	// Goodbye Alice .3
	// Goodbye Bobette .2
	// Goodbye John .1

	names := []string{"Bob", "Alice", "Bobette", "John"}

	for _, n := range names {
		fmt.Printf("Hello %v\n", n)
		fmt.Printf("Goodbye %v\n", n)
	}
}
