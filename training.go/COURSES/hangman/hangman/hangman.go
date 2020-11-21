package hangman

import (
	"fmt"
	"strings"
)

// Game Structure
type Game struct {
	State        string   // Game state
	Letters      []string // Letters in the word to find (type slice strings) [ lettres a trouver]
	FoundLetters []string // Good guesses
	UsedLetters  []string // Used letters
	TurnsLeft    int      // Remaining attempts (nombre de tours restant)
}

// CONSTRUCTOR
// Retourne une instance de jeu (en paramettreon a "turns" (le nombre de tours), "word" (le mot a deviner))
// On retourne un pointeur de Game
func New(turns int, word string) (*Game, error) {

	// on retourne nil et une error formatter
	if len(word) < 3 {
		return nil, fmt.Errorf("Word '%s' must be at least 3 characters. got=%v", word, len(word))
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters)) // construit un slice

	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{}, // slice vide
		TurnsLeft:    turns,
	}

	return g, nil
}

// METHODS
// Createur de Lettre guess - Lettre proposee par le joueur
func (g *Game) MakeAGuess(guess string) {

	guess = strings.ToUpper(guess)

	switch g.State {
	case "won", "lost":
		return
	}

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}

	} else {
		g.State = "badGuess"

		// Decrementation du nombres de tours
		g.LoseTurn(guess)

		// Si on a perdu
		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
}

// Revision des lettres
func (g *Game) RevealLetter(guess string) {

	// On parcourt toutes les lettres
	g.UsedLetters = append(g.UsedLetters, guess)

	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

// Decrement le nombre de tours
// g *Game - pointeur receiver
func (g *Game) LoseTurn(guess string) {
	// g.TurnsLeft -= 1 or g.TurnsLeft = g.TurnsLeft - 1
	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)
}

// Determine if has won (Si on a gagne)
func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}

	return true
}

// Determine if the letter in the Word
func letterInWord(guess string, letters []string) bool {

	// On parcourt les lettres
	for _, l := range letters {

		// Si la lettre est dans le mot on retourne vraie si non false
		if l == guess {
			return true
		}
	}

	return false
}
