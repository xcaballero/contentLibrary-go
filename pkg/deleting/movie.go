package deleting

import (
	"errors"

	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

// ErrMovieNotFound is used when a movie could not be found.
var ErrMovieNotFound = errors.New("movie not found")

func (s *service) DeleteMovie(id string) (listing.Movie, error) {
	movie, err := s.r.GetMovie(id)
	if err != nil {
		return listing.Movie{}, ErrMovieNotFound
	}

	return movie, s.r.DeleteMovie(id)
}
