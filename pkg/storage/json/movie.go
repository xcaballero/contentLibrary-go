package json

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

	newM := Movie{
		ID:        id,
		Title:     m.Title,
		Plot:      m.Plot,
		CreatedAt: time.Now(),
	}

	return newM.toListing(), s.db.Write(CollectionMovie, newM.ID, newM)
}

// GetMovie returns a movie with the specificed ID
func (s *Storage) GetMovie(id string) (listing.Movie, error) {
	var m Movie

	err := s.db.Read(CollectionMovie, id, &m)
	if err != nil {
		return listing.Movie{}, listing.ErrMovieNotFound
	}
	return m.toListing(), nil
}

// GetAllMovies returns all movies
func (s *Storage) GetAllMovies() []listing.Movie {
	list := []listing.Movie{}

	records, err := s.db.ReadAll(CollectionMovie)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, r := range records {
		var m Movie

		if err := json.Unmarshal([]byte(r), &m); err != nil {
			// err handling omitted for simplicity
			return list
		}

		list = append(list, m.toListing())
	}

	return list
}

// DeleteMovie deletes a movie with the specificed ID
func (s *Storage) DeleteMovie(id string) error {
	return s.db.Delete(CollectionMovie, id)
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
