package main

import (
	"log"
	"os"
)

/*
 FILENAME
 log.txt     : nom du fichier

 FLAGS
 os.O_CREATE : creer le fichier s'il n'existe pas
 os.O_APPEND : ecrire et ajouter des informations dans le fichier
 os.O_WRONLY : mode ecriture seulement

 PERMISSION
 0666 (ouvert pour tout le monde)

 =========================================
 $ go build && ./goapprofondi
*/
func main() {

	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("Hello gophers!")
	log.Println("Output in logs.txt")
}
