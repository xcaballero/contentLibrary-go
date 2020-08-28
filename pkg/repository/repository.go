package repository

import (
	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

// Storage provides access to the Storage.
type Storage interface {
	AddMovie(m adding.Movie) (listing.Movie, error)
	GetMovie(id string) (listing.Movie, error)
	GetAllMovies() []listing.Movie
	DeleteMovie(id string) error
}

// Cache provides access to the cache.
type Cache interface {
	AddMovie(m adding.Movie) (listing.Movie, error)
	GetMovie(id string) (listing.Movie, error)
	GetAllMovies() []listing.Movie
	DeleteMovie(id string) error
}

// Repository provides access to the repository.
type Repository interface {
	AddMovie(m adding.Movie) (listing.Movie, error)
	GetMovie(id string) (listing.Movie, error)
	GetAllMovies() []listing.Movie
	DeleteMovie(id string) error
}

type repository struct {
	s Storage
	c Cache
}

// NewRepository returns a new repository
func NewRepository(s Storage, c Cache) Repository {
	return &repository{s, c}
}
