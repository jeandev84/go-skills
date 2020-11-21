package main

import "fmt"

// Entite Movie (Video)
// Renomer la colonne avec `db`
type Movie struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Duration    int    `db:"duration"`
	TrailerURL  string `db:"trailer_url"`
}

// Format to string
func (m Movie) String() string {
	return fmt.Sprintf("id=%v, title=%v, releaseDate=%v, duration=%v, trailerURL=%v",
		m.ID, m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
}
