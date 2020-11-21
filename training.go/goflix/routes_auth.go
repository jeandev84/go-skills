package main

import (
	"fmt"
	"net/http"
)

// handle
func (s *server) handleIndex() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix")
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
