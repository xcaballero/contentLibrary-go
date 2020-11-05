package redis

import (
	"encoding/json"
	"time"

	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
	"github.com/xcaballero/contentLibrary-go/pkg/storage"
)

// Movie defines the storage form of a movie
type Movie struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Plot      string    `json:"plot"`
	CreatedAt time.Time `json:"created_at"`
}

// AddMovie saves the given movie to the repository
func (s *Storage) AddMovie(m adding.Movie) (listing.Movie, error) {
	id, err := storage.GetID("movie")
	if err != nil {
		return listing.Movie{}, err
	}

	movie := Movie{
		ID:        id,
		Title:     m.Title,
		Plot:      m.Plot,
		CreatedAt: time.Now(),
	}

	mString, err := json.Marshal(movie)
	if err != nil {
		return listing.Movie{}, err
	}

	resposta := s.db.Set(s.ctx, movie.ID, mString, 0)

	return movie.toListing(), resposta.Err()
}

// GetMovie returns a movie with the specificed ID
func (s *Storage) GetMovie(id string) (listing.Movie, error) {
	var m Movie

	bMovie, err := s.db.Get(s.ctx, id).Bytes()
	if err != nil {
		return listing.Movie{}, err
	}

	err = json.Unmarshal(bMovie, &m)
	if err != nil {
		return listing.Movie{}, listing.ErrMovieNotFound
	}
	return m.toListing(), nil
}

// GetAllMovies returns all movies
func (s *Storage) GetAllMovies() []listing.Movie {
	list := []listing.Movie{}

	return list
}

// DeleteMovie deletes a movie with the specificed ID
func (s *Storage) DeleteMovie(id string) error {
	return s.db.Del(s.ctx, id).Err()
}

// toListing returns a listing.Movie from a json.Movie
func (m *Movie) toListing() listing.Movie {
	return listing.Movie{
		ID:        m.ID,
		Title:     m.Title,
		Plot:      m.Plot,
		CreatedAt: m.CreatedAt,
	}
}
