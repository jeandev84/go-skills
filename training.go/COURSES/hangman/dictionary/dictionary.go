package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

// Dictionnaire de mots
var words = make([]string, 0, 50)

// Fonction pour charger le fichier
func Load(filename string) error {

	// Ouverture du fichier
	f, err := os.Open(filename)

	if err != nil {
		return err
	}

	// Fermetture du fichier
	defer f.Close()

	// Lecture du fichier line par ligne
	scanner := bufio.NewScanner(f)

	// Tant qu'il y a des choses a scanner
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Renvoie un mot aleatorie (Genere un nombre qui est different)
func PickWord() string {

	// nombre aleatoire
	rand.Seed(time.Now().Unix())

	i := rand.Intn(len(words))

	return words[i]
}
