# INTERFACE

``DEFINITION INTERFACE``
/*
 Une interface est un regroupement nomme  de signatures de fonctions
 Une fonction est un contract

 DEFINITION INTERFACE
 =================================================
type MyInterface interface {
	Foo() error
	Bar() string
}

var a MyInterface = ?

type MyStruct struct {

}

Implementation de methodes de MyInterface par MyStruct il suffit d'ecrire de la facon suivante:

func (m MyStruct) Foo() error {...}
func (m MyStruct) Bar() string {...}

Go deduit automatique que MyStruct est un type de MyInterface

var a MyInterface = &MyStruct{}
err := a.Foo() // appelle l'implementation de MyStruct
*/
func main() {

}


``EVITER DE DEFINIR PLUSIEURS METHODES DANS NOS INTERFACES``
- car cela envoit a un autre programmeur de redefinir plusieurs fonctions
- or il lui interesse qu'une seule
- mieux vaut faire plusieurs petits interfaces pour implementer une une deux methodes


``Ne pas faire ceci``
type MyInterface interface {
    Func1()
    AnotherFunc()
    DidISayILoveFunc()
    AgainAndAgain()
    ICannotStop()
    ThisIsKillingMe()
    EnoughEnougth()
}

``Mais plutot``
- C'est plus composable et facilement utilisable

type MyInterface1 interface {
    Func1()
}


type MyInterface2 interface {
    AnotherFunc()
}


# CONVENTION DES INTERFACES EN GO
``Si une interface a une seule fonction il faut ajouter un suffix -er au nom d'interface``
``cela donne un equivalent de comportement example DO, DOWER``

Example:
package user 

type Saver interface {
    Save(u User) error
}

a l'utilisation
var s user.Saver

# COMPOSER DES INTERFACES

type Saver interface {...}
type Loader interface {...}


type UserSaverLoader interface {
    Saver
    Loader
}