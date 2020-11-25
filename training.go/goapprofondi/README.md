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



# ALIAS EN GO
On utilise parfois des alias pour utiliser des packages ayant le meme nom
import (
    t "text/template"
    h "html/template"
)


func main() {
    t.New("tpl").Parse(`Template TEXT`)
    h.New("tpl").Parse(`Template HTML`)
}


# TYPE ALIAS

type Direction int 

const NORTH Direction = 1
const EAST  Direction = 2
const SOUTH Direction = 3
const WEST  Direction = 4


func blow(dir Direction) {
    fmt.Printf("The wind is blowing %v", dir)
}



// EN JAVA
enums (est une enumeration de valeurs)
enum Level {
    LOW,
    MEDIUM,
    HIGH
}



package main

import "fmt"

// type alias
type ErrorCode int

// definissions des differentes constantes
/*
Alignement de constante:
const ERR_CODE_OK = 0
const ERR_CODE_NOT_FOUND = 1
const ERR_CODE_DUPLICATE = 2
const ERR_CODE_LOCKED = 3
// Groupement de constante
const(
	ERR_CODE_OK = 0
	ERR_CODE_NOT_FOUND = 1
	ERR_CODE_DUPLICATE = 2
	ERR_CODE_LOCKED = 3
)

// Groupement de constante
// iota c'est un mot cle de go qui va incrementer la valeur suvante de 1
// en commencant par 0
const (
	ERR_CODE_OK  ErrorCode = iota // 0
	ERR_CODE_NOT_FOUND // 1
	ERR_CODE_LOCKED   // 2
)
*/

// Groupement de constante
// iota c'est un mot cle de go qui va incrementer la valeur suvante de 1
// en commencant par 0 par default
// ERR_CODE_OK = iota + 1 // ici on le fait commencer par 1
const (
	ERR_CODE_OK ErrorCode = iota
	ERR_CODE_NOT_FOUND
	ERR_CODE_LOCKED
	ERR_CODE_GENERIC
)

// Determine if error code is critical
func (ec ErrorCode) IsCritical() bool {

	return ec == ERR_CODE_LOCKED || ec == ERR_CODE_NOT_FOUND
}

// Print error
func printErrCode(c ErrorCode) {
	fmt.Printf("code=%v, critical=%v\n", c, c.IsCritical())
}

// Entry point
func main() {

	code := ERR_CODE_OK
	printErrCode(code)
}


# CROSS COMPILATION
// Compiler ce programme comme suit
// pour adapter a toute sorte de machine:
// prefixer go build par `` env GOOS=linux GOARCH=amd64 ``

- Linux
`` env GOOS=linux GOARCH=amd64 go build ``

- Windows
`` env GOOS=windows GOARCH=amd64 go build ``