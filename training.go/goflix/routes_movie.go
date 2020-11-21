package main

import (
	"log"
	"net/http"
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
