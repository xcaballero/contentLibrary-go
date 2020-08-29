package deleting

import (
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

// Repository provides access to the repository.
type Repository interface {
	// DeleteMovie deletes the movie with the given ID from the repository.
	DeleteMovie(string) error
	// GetMovie returns the movie with given ID.
	GetMovie(string) (listing.Movie, error)
}

// Service provides adding operations.
type Service interface {
	DeleteMovie(string) (listing.Movie, error)
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}
