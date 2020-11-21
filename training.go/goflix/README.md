# GOFLIX (WEB PROJECT) Structure Mat Ryer

- 1. Creer notre go module

``projects $ pwd``
``projects $ mkdir goflix``
``projects $ cd goflix``
``projects/goflix $ go mod init training.go/goflix``
``to github : projects/goflix $ go mod init github.com/nomutilisateur/goflix``
``afficher le contenu du fichier go.mod $ cat go.mod``

`` Creation du fichier main.go``
package main

import "fmt"

func main() {
	fmt.Println("GoFlix")
}

Executer la commande ``$ go build .``
Il sera generer un fichier gofix
On peut alors lancer le programme de la sorte suivante: ``$ ./goflix``


- 2. Installation packages 
``jmoiron/sqlx`` pour interagir avec la base de donnees https://github.com/jmoiron/sqlx
``mattn/go-sqlite3`` https://github.com/mattn/go-sqlite3

- go get

`` $ go build && ./goflix``

- 3. Installation du package `` gorilla/mux `` pour le routing ``https://github.com/gorilla/mux``


`` $ go get github.com/gorilla/mux ``



