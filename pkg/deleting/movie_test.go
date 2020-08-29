package deleting_test

import (
	"testing"
	"time"

	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/deleting"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
	"github.com/xcaballero/contentLibrary-go/pkg/storage"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteMovie(t *testing.T) {
	m1 := adding.Movie{
		Title: "Test Movie 1",
		Plot:  "Plot One",
	}

	mS := new(mockStorage)

	dS := deleting.NewService(mS)

	movie1, err := mS.AddMovie(m1)
	require.NoError(t, err)
	assert.EqualValues(t, m1.Title, movie1.Title)
	assert.EqualValues(t, m1.Plot, movie1.Plot)

	movie1, err = dS.DeleteMovie(movie1.ID)
	require.NoError(t, err)
	assert.EqualValues(t, m1.Title, movie1.Title)
	assert.EqualValues(t, m1.Plot, movie1.Plot)

	_, err = dS.DeleteMovie("wrong_id")
	require.Error(t, err)

	movies := mS.GetAllMovies()
	assert.Len(t, movies, 0)
}

type mockStorage struct {
	movies []listing.Movie
}

// DeleteMovie deletes a movie with the specified ID
func (mS *mockStorage) DeleteMovie(id string) error {
	for i, m := range mS.movies {
		if m.ID == id {
			if len(mS.movies) == 1 {
				mS.movies = []listing.Movie{}
				return nil
			}
			copy(mS.movies[i:], mS.movies[i+1:])          // Shift a[i+1:] left one index.
			mS.movies[len(mS.movies)-1] = listing.Movie{} // Erase last element (write zero value).
			mS.movies = mS.movies[:len(mS.movies)-1]
			return nil
		}
	}

	return listing.ErrMovieNotFound
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

func (mS *mockStorage) AddMovie(m adding.Movie) (listing.Movie, error) {
	id, err := storage.GetID("movie")
	if err != nil {
		return listing.Movie{}, err
	}

	movie := listing.Movie{
		ID:        id,
		Title:     m.Title,
		Plot:      m.Plot,
		CreatedAt: time.Now(),
	}
	mS.movies = append(mS.movies, movie)

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
