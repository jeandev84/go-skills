# Struct (Structure)


``Un Struct permet de regrouper des variables dans un type personnalise``
Regrouper des donnees et des functions dans un package donnee

``type <NomStruct> struct {``
  ``var1 int``
  ``var2 string``
  ``var3 float64`
``}``

``Example : ``
``type User struct {``
  ``Name  string``
  ``Email string``
  ``Age   int``
``}``

NB en Go:
- Majuscule pour mettre en public
- Minuscule pour mettre en private

* Un struct ne peut contenir que des variables
* En go on ne definit pas de methodes et de fonctions a l'interieure d'un struct
* La rege de la visibilite de package s'applique pour:
  - Le Struct lui meme
  - les variables de la struct


# Les Fonctions
``Effectue des traitements donnees``


# Les Methodes 
``Sont des fonctions qui permettent d'effectuer des traitements de donnees, mais elle vont directement interagir avec les Structs``


Example:

type User struct {
    Name string
}


// Methodes
func (u User) SayHello() {
   fmt.Printf("Hello %v!\n", u.Name)
}