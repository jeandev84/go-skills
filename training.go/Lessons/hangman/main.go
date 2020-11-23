package main

import (
	"fmt"
	"os"

	"training.go/hangman/dictionary"
	"training.go/hangman/hangman"
)

// Entry point of project
func main() {

	// Chargement du fichier de dictionnaire
	err := dictionary.Load("words.txt")

	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		os.Exit(1)
	}

	// Pointeur de Game
	// g := hangman.New(8, "Golang")
	g := hangman.New(8, dictionary.PickWord())

	// Welcome
	hangman.DrawWelcome()

	// Etat du JEU
	guess := ""

	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost": // si g.State == won ou lost
			os.Exit(0) // code erreur 0, pour dire que tout s'est bien passe
		}

		// l (lettre saisi), err (error)
		l, err := hangman.ReadGuess()

		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1) // sortie du program avec un code d'erreur qui dit ca n'a pas marche
		}

		guess = l
		g.MakeAGuess(guess)
	}

}
