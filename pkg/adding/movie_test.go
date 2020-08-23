package adding_test

import (
	"testing"
	"time"

	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
	"github.com/xcaballero/contentLibrary-go/pkg/storage"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddMovie(t *testing.T) {
	m1 := adding.Movie{
		Title: "Test Movie 1",
		Plot:  "Plot One",
	}

	m2 := adding.Movie{
		Title: "Test Movie 2",
		Plot:  "Plot Two",
	}

	mR := new(mockStorage)

	s := adding.NewService(mR)

	movie1, err := s.AddMovie(m1)
	require.NoError(t, err)
	assert.EqualValues(t, m1.Title, movie1.Title)
	assert.EqualValues(t, m1.Plot, movie1.Plot)

	movie1, err = s.AddMovie(m1)
	require.Error(t, err)
	assert.EqualValues(t, listing.Movie{}, movie1)

	movie2, err := s.AddMovie(m2)
	require.NoError(t, err)
	assert.EqualValues(t, m2.Title, movie2.Title)
	assert.EqualValues(t, m2.Plot, movie2.Plot)

	movies := mR.GetAllMovies()
	assert.Len(t, movies, 2)
}

type mockStorage struct {
	movies []adding.Movie
}

func (mS *mockStorage) AddMovie(m adding.Movie) (listing.Movie, error) {
	mS.movies = append(mS.movies, m)

	id, err := storage.GetID("movie")
	if err != nil {
		return listing.Movie{}, err
	}

	return listing.Movie{
		ID:        id,
		Title:     m.Title,
		Plot:      m.Plot,
		CreatedAt: time.Now(),
	}, nil
}

func (mS *mockStorage) GetAllMovies() []listing.Movie {
	movies := []listing.Movie{}

	for _, mm := range mS.movies {
		m := listing.Movie{
			Title: mm.Title,
			Plot:  mm.Plot,
		}
		movies = append(movies, m)
	}

	return movies
}
