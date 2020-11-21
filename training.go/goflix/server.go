package main

// struct server (application)
type server struct {
	store Store
}

// return une nouvelle instance de server
func newServer() *server {

	s := &server{}

	return s
}
