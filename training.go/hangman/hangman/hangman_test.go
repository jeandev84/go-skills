package hangman

import (
	"testing"
)

func TestLetterInWord(t *testing.T) {

	word := []string{"b", "o", "b"}
	guess := "b"
	hasLetter := letterInWord(guess, word)

	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {

	word := []string{"b", "o", "b"}
	guess := "a"
	hasLetter := letterInWord(guess, word)

	if hasLetter {
		t.Errorf("Word %s does not contain letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {

	// On obtient seulement l'erreur, le Game ne nous interesse pas
	_, err := New(3, "")

	if err == nil {
		t.Errorf("Error should be returned when using an invalid word=''")
	}
}

// Test sur une bonne supposition
func TestGameGoodGuess(t *testing.T) {

	g, _ := New(3, "bob")
	g.MakeAGuess("b")

	validState(t, "goodGuess", g.State)
}

func TestGameAlreadyGuessed(t *testing.T) {

	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("b")

	validState(t, "alreadyGuessed", g.State)
}

func TestGameWon(t *testing.T) {

	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	g.MakeAGuess("b")

	validState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {

	g, _ := New(3, "bob")
	g.MakeAGuess("a")
	g.MakeAGuess("z")
	g.MakeAGuess("e")

	validState(t, "lost", g.State)
}

// Fonction privee
func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("state should be '%v'. got=%v", expectedState, actualState)
		return false
	}

	return true
}
