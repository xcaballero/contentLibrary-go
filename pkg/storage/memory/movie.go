package memory

import (
	"time"

	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
	"github.com/xcaballero/contentLibrary-go/pkg/storage"
)

// Movie defines the storage form of a movie
type Movie struct {
	ID        string
	Title     string
	Plot      string
	CreatedAt time.Time
}

// AddMovie saves the goven movie to the repository
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
	s.movies = append(s.movies, newM)

	return newM.toListing(), nil
}

// GetMovie returns a movie with the specified ID
func (s *Storage) GetMovie(id string) (listing.Movie, error) {
	for _, m := range s.movies {
		if m.ID == id {
			return m.toListing(), nil
		}
	}

	return listing.Movie{}, listing.ErrMovieNotFound
}

// GetAllMovies return all movies
func (s *Storage) GetAllMovies() []listing.Movie {
	var movies []listing.Movie

	for _, m := range s.movies {
		movies = append(movies, m.toListing())
	}

	return movies
}

// DeleteMovie deletes a movie with the specified ID
func (s *Storage) DeleteMovie(id string) error {
	for i, m := range s.movies {
		if m.ID == id {
			if len(s.movies) == 0 {
				return listing.ErrMovieNotFound
			} else if len(s.movies) == 1 {
				s.movies = []Movie{}
				return nil
			} else {
				copy(s.movies[i:], s.movies[i+1:])  // Shift a[i+1:] left one index.
				s.movies[len(s.movies)-1] = Movie{} // Erase last element (write zero value).
				s.movies = s.movies[:len(s.movies)-1]
				return nil
			}
		}
	}

	return listing.ErrMovieNotFound
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
