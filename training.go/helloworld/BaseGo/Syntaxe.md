######## BASE GO ############

# Types
bool  `true / false`
string ` "Hello", "Goodbye"`
int    `10, 20, 42`
byte   `01001110 ==> storage sur un octet (8 bits)`
float32 float64  `3.14 ==> equivalent a float et double`

========================================
# LISTING ALL TYPES
#======================

# Boolean
bool

# Chaine de characteres
string

# Representer des entiers en bytes en fonction de la capacite de l'ordinateur
int int8 int16 int32 int64
uint uint8 unit16 uint32 unit64 uinitptr

byte  `alias for unit8`

# rune pour generer un Unicode (generer les differentes langues qui existent dans le monde)
rune `alias for int32 represents a Unicode code point`


# Generer les nombres a virgules
float32 float64 


# Representer les nombres mathematiques complexes
complex64 complex128


# SYNTAX Declaration de variable

var nameOfVariable Type = value
 
Example:

 var age int = 20
 var age int // age = 0 (par default)


 var name string = "Bob"
 var name string // name = "" (par defaut)


 Visibilite des variables en GO
 ===========================================
 - Variable en minuscule (Variable prive)
 - Variable en majuscule (Variable public)

 exmaple:
 name := "Hello" (variable prive) cette variable sera accessible seulement a l'interieur de mon package
 Name := "Hello" (variable public) cette variable sera accessible a l'exterieur de mon package


 Flux : If ... else
 ===========================================

 if <condition> {
    // code
 }


age := 15

if age > 10 {
    // Bob is old enough
}



if age > 10 {
    // code
} else {
    // condition age <= 10
    // alternative code
}


if age > 10 {
    // code

} else if age > 5 {
    // condition age <= 10 ET age > 5
    // alternative 1 code
} else if age > 2 {
    // condition age <= 10 ET age > 5 ET age > 2
    // alternative 1 code
} else {
    // condition age <= 10
    // alternative code
}



if <condition1> ET <condition2> {
    // code
}


if age > 10 && age < 15 {
   // code
}


/*
 Comparaisons
 Egal ==
 Different !=
 Strictement inferieur <
 Strictement superieur >
 Inferieur ou egal <=
 Superieur ou egal >=
*/
 

 (ET)  && - les 2 condtions sont vraies
 (OU)  || - Au moins 1 est vraie
 (NON) !  - La condition doit etre false


 condition 1, condition 2, resultat
 0               0           0
 0               1           0
 1               0           0
 1               1           1

 

 SWITCH
===========================================================
switch <variable>

case <cas 1>:
   // code cas 1
case <cas 2>:
   // code cas 2
default:
   // cas par default (facultatif)


switch age {
    case 10:
       // code cas 10
    case 20:
       // code cas 20
    default:
       // code default
}


Operations impossibles
===================================
ten := "10"
var x int = ten + 1 // erreur de compilation

var n int = 1
n = n + 2.14 // erreur de compilation

Convertion
var f float64 = 2.0
var n int = int(f)



Functions
====================================

func <NomFonction>(<param1> type, <param2> type) <TypeRetour> {
    // code
}


nom de function en minuscule est prive
nom de function en majuscule est publique


func greeting() {
    fmt.Println("Hello there")
}

greeting() // output: Hello there


func greeting(name string) {
    fmt.Printf("Hello there %s\n", name)
}


greeting("Bob") // output: Hello there Bob



func oldEnough(age int) bool {
    return age >= 18
}


canPlay := oldEnough(15) // returns false


Fonction Multiple
======================================

func MyFunc() (type1, type2) {
    return val1, val2
}


Cette fonction renvoie 2 types
func MultReturn() (int, string) {
    return 4, "hello"
}

Appeler la fonction
MultReturn()


Recuperation des valeurs retournees par la fonction
num, name := MultReturn()


Recuperation de la valeur de retour mais en ignorant la valeur int
_, name := MultReturn()



TABLEAU DEFINTION SYNTAXE
===================================================

Voici comment declarer un tableau en Go!

var nom[taille]type

Example:
- Declaration un tableau de nom "t" de taille 5 de nombres entiers
var t[5]int 

- Affectation (Set)
t[3] = 12

- Sortie: 
output: 12 ===> Lecture (Get)
fmt.Println("", t[3])

- Recuperer la taille d'un tableau:
output: 5 ==> Taille de tableau
len(t)


SLICES (Array de taille dynamique)
======================================
un slice va pouvoir augmenter en fonction du contenu

SLICE = Tranche
Un Slice represente une tranche d'un tableau

Syntaxe:

s := make([]type, taille, capacite)
. Taille: correspond au nombre d'elements du slice
. Capacite[facultatif]: nombre d'elements du tableau par default il prendra la taille du tableau