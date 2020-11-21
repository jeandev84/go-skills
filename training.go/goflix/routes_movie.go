package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Struct pour le format JSON de movies
type jsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerURL  string `json:"trailer_url"`
}

// movie list handle
func (s *server) handleMovieList() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// recuperer la liste des movies
		movies, err := s.store.GetMovies()

		if err != nil {

			log.Printf("Cannot load movies. err=%v\n", err)
			// TODO handle response to the client
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		// conversion des movies en JSON Format
		// la reponse (resp) sera un tableau de json movie
		var resp = make([]jsonMovie, len(movies))

		for i, m := range movies {
			resp[i] = mapMovieToJson(m)
		}

		// TODO response JSON Format for writer
		s.respond(w, r, resp, http.StatusOK)
	}

}

// Fonction pour voir les details du movie (video)
func (s *server) handleMovieDetail() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Recuperation des valeurs venant de la requette example "id"
		vars := mux.Vars(r)

		// 10 est la base octet, binaire, decimal, hexadecimal ...
		id, err := strconv.ParseInt(vars["id"], 10, 64)

		if err != nil {
			log.Printf("Cannot parse id to int. err=%v\n", err)
			s.respond(w, r, nil, http.StatusBadRequest)
			return
		}

		m, err := s.store.GetMovieById(id)

		if err != nil {
			log.Printf("Cannot load movie. err=%v\n", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapMovieToJson(m)
		s.respond(w, r, resp, http.StatusOK)
	}
}

// Create un movie
/*
{
	 "title": "The Avengers",
	 "release_date": "2012-01-04",
	 "duration": 143,
	 "trailer_url": "https://www.youtube.com/watch?v=s0rN10pCv1"
}
*/
func (s *server) handleMovieCreate() http.HandlerFunc {

	// request body or content
	type request struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerURL  string `json:"trailer_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := request{}
		err := s.decode(w, r, &req)

		if err != nil {
			log.Printf("Cannot parse movie body. err=%v", err)
			s.respond(w, r, nil, http.StatusBadRequest)
			return
		}

		// Create a Movie
		m := &Movie{
			ID:          0,
			Title:       req.Title,
			ReleaseDate: req.ReleaseDate,
			Duration:    req.Duration,
			TrailerURL:  req.TrailerURL,
		}

		// Store the Movie in the DB
		err = s.store.CreateMovie(m)

		if err != nil {
			log.Printf("Cannot create movie in DB. err=%v\n", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapMovieToJson(m)
		s.respond(w, r, resp, http.StatusOK)
	}
}

// JSON Movie format JSON
func mapMovieToJson(m *Movie) jsonMovie {

	return jsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerURL:  m.TrailerURL,
	}
}
