# PROJECT JEU DU PENDU 
``Fonctionnalites``

- Affichage & le Dessin du pendu en console
- Chargement aleatoire du mot a partir d'un fichier de dictionnaire
- Affichage de l'etat courant de la partie au cours du jeu
  * avec le pendu
  * les lettres trouvees qui permettent de composer le mot
  * les lettres qui ont ete deja utilisees pour aider le joueur


# Architecture

* Forte separation logique de jeu /io (input/output  - entree et sortie du programme)
 - Input clavier (Saisi du clavier)
 - Output ecran console (L'affichage console)

* Chargement dictionnaire isole
 - Package (hangman) qui contiendra la logique du jeu biensur
 - Package (dictionary) pour avoir le dictionnaire des mots
 - main (installer tout le cablage qui fait discuter tous les 3 packages ensemble)

* Tests unitaires en Go


# Generateur ASCII ART (Pour generer des tags)
``http://patorjk.com/software/taag/``

``  |\     /|(  ___  )( (    /|(  ____ \(       )(  ___  )( (    /| ``
``	| )   ( || (   ) ||  \  ( || (    \/| () () || (   ) ||  \  ( | ``
``	| (___) || (___) ||   \ | || |      | || || || (___) ||   \ | | ``
``	|  ___  ||  ___  || (\ \) || | ____ | |(_)| ||  ___  || (\ \) | ``
``	| (   ) || (   ) || | \   || | \_  )| |   | || (   ) || | \   | ``
``	| )   ( || )   ( || )  \  || (___) || )   ( || )   ( || )  \  | ``
``	|/     \||/     \||/    )_)(_______)|/     \||/     \||/    )_) ``

# LUNCH PROGRAM
``$ go run main.go``


# TEST GO (Norme et Convention a respecter pour effectuer des tests en go)
Caracteristiques d'une fonction de test
(La convention en GO est de nomme le nom des methodes par TestXxx()
Unique parametre t *testing.T)

- Nom de fonction de type TestXxx()
- Unique parametre t *testing.T
- Appelle t.Error / t.Failure pourque le test echoue
- Le fichier doit s'appeler nom-du-fichier-a-tester_test.go (avec le suffix _test.go)

* Executer un test en GO :$ go test ./hangman (hangman package name)
``$ go test ./# Dans le project, et tous les tests que go aurait rencontre seront executes``

``hangman$ go test ./hangman``
``ok      training.go/hangman/hangman     0.003s``

``mymath/logic.go``

package mymath


func Mult(x int, y int) int {
   return x * y
}


six := Mult(2, 3)

``mymath/logic_test.go``

import "testing"


func TestMult(t *testing.T) {
    total := Mult(2, 3)

    if total != 6 {
        t.Errorf("Incorrect Mult. got=%d, want=%d", total, 6)
    }
}