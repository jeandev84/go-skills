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



- 4. Middlewares

Un middleware est un bout de code qu'on va inserer juste avant l'execution d'une requette ou juste apres l'execution de la generation d'une reponse

Example on veut faire un log a chaque fois qu'un client execute une reponse


- 5. JWT (JSON WEB TOKEN) `` https://jwt.io/ `` pour securiser des api

Package go : `` github.com/dgrijalva/jwt-go ``


# LUNCH GOFLIX WEB APPLICATION

`` $ go build && ./goflix ``

Authorization
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
)

// handle
func (s *server) handleIndex() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix")
	}
}

// Create JWT Token4
/*
{
	"username": "golang",
	"password": "rocks"
}

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDYwMDk5MjIsImlhdCI6MTYwNjAwNjMyMiwidXNlcm5hbWUiOiJnb2xhbmcifQ.EY3AmMnYO_Qqf8wBrFF0ZX6cqSGbLyMe-rJC3QdqO6Y"
}
*/
func (s *server) handleTokenCreate() http.HandlerFunc {

	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type response struct {
		Token string `json:"token"`
	}

	type responseError struct {
		Error string `json:"error"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		// Parsing login body
		req := request{}
		err := s.decode(w, r, &req)

		if err != nil {

			msg := fmt.Sprintf("Cannot parse login body. err=%v", err)

			log.Println(msg)

			s.respond(w, r, responseError{
				Error: msg,
			}, http.StatusBadRequest)

			return
		}

		// Check credentials
		found, err := s.store.FindUser(req.Username, req.Password)

		if err != nil {

			msg := fmt.Sprintf("Cannot find user. err=%v", err)
			s.respond(w, r, responseError{
				Error: msg,
			}, http.StatusInternalServerError)

			return
		}

		/*
			if req.Username != "golang" || req.Password != "rocks" {

				s.respond(w, r, responseError{
					Error: "Invalid credentials",
				}, http.StatusUnauthorized)

				return
			}
		*/

		if !found {

			s.respond(w, r, responseError{
				Error: "Invalid credentials",
			}, http.StatusUnauthorized)

			return
		}

		/* On a pu recuper les informations de l'utilisateur on peut creer un TOKEN
		// (algodecodage, les references a coder)
		// username: utilisateur, exp : date expiration token, iat: issue at time (date a laquelle le token a ete fournit au client)
		*/

		// Generate Token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": req.Username,
			"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			"iat":      time.Now().Unix(),
		})

		tokenStr, err := token.SignedString([]byte(JWT_APP_KEY))

		if err != nil {

			msg := fmt.Sprintf("Cannot generate JWT. err=%v", err)

			log.Println(msg)

			s.respond(w, r, responseError{
				Error: msg,
			}, http.StatusInternalServerError)

			return
		}

		// Si tout va bien on envoie le token en reponse
		s.respond(w, r, response{
			Token: tokenStr,
		}, http.StatusOK)

	}
}

/*
Interface interne
type Handler interface {
   ServeHTTP(ResponseWriter, *Request)
}

func (s *server) handleIndex() http.Handler {
   return implementation
}
*/


# ASPECT POUR ALLEZ PLUS LOIN DANS LE PROJECT GOFLIX
- 1. Creer un system de pagination 
- 2. Complexifier la BDD (genres, acteurs, realisateurs, notes) Movie
- 3. Ajout d'une liste de souhaits qui sera associe a un utilisateur
     Et donc les routes pour modifier cette liste
- 4. Faire beaucoup tests http, routing
