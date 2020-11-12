package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Lecture d'une entree au clavier (Entrre standard est le clavier : os.Stdin)
// Sortie standard est la console
var reader = bufio.NewReader(os.Stdin)

// Lecture d'une supposition
// Retour guess (saisi), err error
func ReadGuess() (guess string, err error) {

	valid := false

	for !valid {

		fmt.Print("What is your letter? ")

		// Il fait sa lecture jusqu'a ce que le joueur aurait tape sur Entrez
		guess, err = reader.ReadString('\n')

		// early return (traitement d'erreur en premier)
		if err != nil {
			return guess, err
		}

		// enlever les espaces
		guess = strings.TrimSpace(guess)

		if len(guess) != 1 {
			fmt.Printf("Invalid letter input. letter=%v, len=%v\n", guess, len(guess))
			continue
		}

		valid = true
	}

	return guess, nil
}
