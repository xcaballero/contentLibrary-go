package adding

import (
	"errors"

	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

// ErrDuplicateMovie is used when a movie already exists.
var ErrDuplicateMovie = errors.New("movie already exists")

// Movie defines the properties of a movie to be added
type Movie struct {
	Title string `json:"title"`
	Plot  string `json:"plot"`
}

// AddMovie persists the given movie(s) to storage
func (s *service) AddMovie(m Movie) (listing.Movie, error) {
	// make sure we don't add any duplicates
	existingMovies := s.r.GetAllMovies()
	for _, e := range existingMovies {
		if m.Title == e.Title &&
			m.Plot == e.Plot {
			return listing.Movie{}, ErrDuplicateMovie
		}
	}

	movie, err := s.r.AddMovie(m) // error handling omitted for simplicity

	return movie, err
}
