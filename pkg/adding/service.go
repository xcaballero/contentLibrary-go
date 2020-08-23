package adding

import (
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

// Repository provides access to the repository.
type Repository interface {
	// AddMovie saves a gicen movie to the repository.
	AddMovie(Movie) (listing.Movie, error)
	// GetAllMovies returns all movies saved in storage.
	GetAllMovies() []listing.Movie
}

// Service provides adding operations.
type Service interface {
	AddMovie(Movie) (listing.Movie, error)
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}
