package main

import (
	"fmt"

	"training.go/helloworld/player"
)

// Entry point of Application
func main() {

	// Method 1: Initialiser un struct avec arguments, il faut absolument passer toutes les valeurs
	// var p1 player.Player = player.Player{"Bob", 10, {"http://avatar.png"}}

	// Method 2: Intialiser un struct sans passer d'arguments
	var p1 player.Player

	p1.Name = "Bob"
	p1.Age = 10

	fmt.Printf("Player 1: %v\n", p1)
	fmt.Printf("p1.Name=%v, p1.Age=%v\n", p1.Name, p1.Age)

	// Method3: Initialiser un avatar du player en passant directement les champs
	// Typage et affectation de valeur
	avatar := player.Avatar{"http://avatar.jpg"}
	fmt.Printf("avatar : %v\n", avatar)

	// Method 4: Type et assignation des valeurs sous forme de cle et valeur
	// Methode recommandee
	// car cela nous donne le choix de choisir les champs qu'on veut assigner une valeur
	// La convention en go il faut mettre toujours une virgule quand on passw a la ligne
	p3 := player.Player{
		Name: "John",
		Age:  25,
		Avatar: player.Avatar{
			Url: "http://url.com",
		},
	}

	fmt.Printf("Player 3 : %v\n", p3)

	// Recommande
	p4 := player.New("Bobette")
	p4.Avatar = avatar
	fmt.Printf("Player 4 : %v\n", p4)
}
