package listing_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

func TestGetMovie(t *testing.T) {
	m1 := listing.Movie{
		ID:        "movie_1",
		Title:     "Test Movie 1",
		Plot:      "Plot One",
		CreatedAt: time.Now(),
	}

	m2 := listing.Movie{
		ID:        "movie_2",
		Title:     "Test Movie 2",
		Plot:      "Plot Two",
		CreatedAt: time.Now(),
	}

	mR := new(mockStorage)
	mR.movies = append(mR.movies, m1)
	mR.movies = append(mR.movies, m2)

	s := listing.NewService(mR)

	movie1, err := s.GetMovie(m1.ID)
	require.NoError(t, err)
	assert.EqualValues(t, m1.Title, movie1.Title)
	assert.EqualValues(t, m1.Plot, movie1.Plot)

	movie2, err := s.GetMovie(m2.ID)
	require.NoError(t, err)
	assert.EqualValues(t, m2.Title, movie2.Title)
	assert.EqualValues(t, m2.Plot, movie2.Plot)
}

func TestListMovies(t *testing.T) {
	m1 := listing.Movie{
		ID:        "movie_1",
		Title:     "Test Movie 1",
		Plot:      "Plot One",
		CreatedAt: time.Now(),
	}

	m2 := listing.Movie{
		ID:        "movie_2",
		Title:     "Test Movie 2",
		Plot:      "Plot Two",
		CreatedAt: time.Now(),
	}

	mR := new(mockStorage)
	mR.movies = append(mR.movies, m1)
	mR.movies = append(mR.movies, m2)

	s := listing.NewService(mR)

	movies := s.ListMovies()
	assert.Len(t, movies, 2)
	assert.EqualValues(t, m1, movies[0])
	assert.EqualValues(t, m2, movies[1])
}

type mockStorage struct {
	movies []listing.Movie
}

func (mS *mockStorage) GetMovie(id string) (listing.Movie, error) {
	movie := listing.Movie{}

	for _, m := range mS.movies {
		if m.ID == id {
			return m, nil
		}
	}

	return movie, listing.ErrMovieNotFound
}

func (mS *mockStorage) GetAllMovies() []listing.Movie {
	movies := []listing.Movie{}

	for _, m := range mS.movies {
		movies = append(movies, m)
	}

	return movies
}
