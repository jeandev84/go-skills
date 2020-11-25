# SYNTAXE CONSTANTE

`` const refresh_day = 500 ``  // un nombre entier
`` const pi = 3.145926 ``      // un nombre a virgule
`` const secret_key = '0123456789abcdef' `` // une string

- Impossible d'assigner une nouvelle valeur a une constante comme ceci :

`` const refresh_day = 500 ``
`` refresh_delay = 2 !!! erreur de compilation ``


- const a la compilation 
`` // Ce que l'on ecrit ``
`` const reflesj_delay = 500 ``
`` delay := 2 ``
`` if delay >= refresh_delay { ... } ``

`` // Ce que Go genere a la compilation ``
`` delay := 2 ``
`` if delay >= 500 { ... } // Le const est remplace par la valeur ``


# Example

package main 


import "log"


func main() {
    log.Println("Hello gophers!")
}


2020/06/31 16:42:23 Hello gophers!


# Bibliotheque recommandee pour les logs:

`` https://github.com/Siruspen/logrus ``


# Fonction litterale ou fonction anonyme comme les fonctions de Closure
func main() {
    
   La fonction est definie sans nom
   les () qui suivent indiquent l'appel a la fonction
   func() {
       fmt.Println("Hello!")
   }()
}


# High Order Function (Fonction du premier ordre)

`` func Slice(slice interface{}, less func(i, j int) bool) ``
- less : est un parametre de type une fonction
- il prend 2 entiers i et j en parametre
- il renvoie un bool


# Stuct anonyme
package main

import (
	"fmt"
)


func main() {
	
	// Struct Anonyme (pas de nom)
	s := struct{ // Declaration du struct
		name: string,
		grade: int
	}{
	   "Bob",
	   70
	}

	fmt.Println(s)
}


# TABLE TEST

func inc(i int) int {
    return i + 1
}


func TestInc1(t *testing.T) {
    v := inc(1)
    if v != 2 {
        t.Errorf("Invalid inc(1) result. expected=%v, got=%v", 2, 1)
    }
}


func TestInc2(t *testing.T) {
    v := inc(2)
    if v != 3 {
        t.Errorf("Invalid inc(2) result. expected=%v, got=%v", 3, 1)
    }
}



func TestIncTable(t *testing.T) {

    var tests = []struct {
        in int  // argument d'entree
        expected int // resultat attendu
    }{
        {1, 2},    // definition des cases
        {2, 3},
        {3, 4},
    }
    
    // tt : dans le sens de "test table"
    for _, tt := range tests {
        v := inc(tt.in)
        if v != tt.expected {
            t.ErrorF("Invalid inc(%v) result. expected=%v, got=%v", tt.in, tt.expected, v)
        }
    }
}


# LANCER TEST 
`` go test . ``

# BenchmarkPrint


// La fonction commence par BenchmarkXxxx
// Elle prend en parametre un type testing.B
func BenchmarkPrint(b *testing.B) {
    
    // La boucle pour executer le Benchmark
    for i := 0; i < b.N; i++ {
        
        // Le corps de ce que l'on veut mesurer
        fmt.Sprintf("Hi Gophers")
    }
}

Excecuter Benchmark:
. dossier courant
`` go test -bench=. ``



# ERROR CUSTOM
 En Go le type "error" n'est qu'une interface renvoyant un message d'erreur
type error interface {
    Error() string
}


type MyError struct {
    Module string
    Err error
}


func (e *MyError) Error() string {

    return fmt.Sprintf("Module=%v, error=%v", e.Module, e.Err)
}


Utilisation de l'erreur custom

func process() error {
    return &MyError{
        Module: "main",
        Err: fmt.Errorf("Could not execute process()")
    }
}



func main() {
    
    err := process()

    if err != nil {
        // Affiche Module=main, error=Cound not execute process()
        fmt.Println(err)
    }
}