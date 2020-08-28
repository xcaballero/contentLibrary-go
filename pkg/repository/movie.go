package repository

import (
	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

func (r *repository) AddMovie(m adding.Movie) (listing.Movie, error) {
	return r.s.AddMovie(m)
}

func (r *repository) GetMovie(id string) (listing.Movie, error) {
	return r.s.GetMovie(id)
}

func (r *repository) GetAllMovies() []listing.Movie {
	return r.s.GetAllMovies()
}

func (r *repository) DeleteMovie(id string) error {
	return r.s.DeleteMovie(id)
}
