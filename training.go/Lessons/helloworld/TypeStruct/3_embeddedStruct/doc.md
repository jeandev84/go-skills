# Embedded Struct

``Identifie le lien entre 2 Structs``


Quel est le lien qui est exprime entre les 2 structs

type Address struct {
     City string
}


type User struct {
    Name string
    Addr Address
}


Relation de possession Owner Ship
- un User a une Address (OneToOne)

Un XXX est un YYY (Heritage)


Go prefere la composition avec l'Embedded Struct
Embedded Struct signifit on va inclure un Struct a l'interieur d'un autre
Ici Address est inclus dans User et on peut l'utiliser
Address est embarque dans le type User (On peut dire User est une Address)

type Address struct {
     City string
}


type User struct {
    Name string
    Address // pas nom de variables
}

var u User
u.City = "Londres" // City est directement accessible