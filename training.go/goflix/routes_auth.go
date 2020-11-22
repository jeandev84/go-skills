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
