package listing

import (
	"errors"
	"time"
)

// ErrMovieNotFound is used when a movie could not be found.
var ErrMovieNotFound = errors.New("movie not found")

// Movie defines the properties of a movie to be listed
type Movie struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Plot      string    `json:"plot"`
	CreatedAt time.Time `json:"created_at"`
}

// GetMovies returns all movies
func (s *service) GetMovies() []Movie {
	return s.r.GetAllMovies()
}

// GetMovie returns a movie
func (s *service) GetMovie(id string) (Movie, error) {
	return s.r.GetMovie(id)
}
