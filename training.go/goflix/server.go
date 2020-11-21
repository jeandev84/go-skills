package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// struct server (application)
type server struct {
	router *mux.Router
	store  Store
}

// return une nouvelle instance de server
func newServer() *server {

	// create new server or application
	s := &server{
		router: mux.NewRouter(),
	}

	// load routes
	s.routes()

	// return server or application
	return s
}

// Delegation HTTP request / response
func (s *server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Response
func (s *server) respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}
