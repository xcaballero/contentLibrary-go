package json_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/storage/json"
)

func TestAddMovie(t *testing.T) {
	s, err := json.NewStorage()
	require.NoError(t, err)

	newM := adding.Movie{
		Title: "New Movie",
		Plot:  "New movies plot.",
	}

	m, err := s.AddMovie(newM)

	require.NoError(t, err)
	assert.EqualValues(t, newM.Title, m.Title)
	assert.EqualValues(t, newM.Plot, m.Plot)

	require.NoError(t, s.DeleteMovie(m.ID))
}

func TestGetMovie(t *testing.T) {
	s, err := json.NewStorage()
	require.NoError(t, err)

	newM := adding.Movie{
		Title: "New Movie",
		Plot:  "New movies plot.",
	}

	m, err := s.AddMovie(newM)
	require.NoError(t, err)

	getM, err := s.GetMovie(m.ID)
	require.NoError(t, err)
	assert.ObjectsAreEqual(m, getM)

	_, err = s.GetMovie("wrong_id")
	require.Error(t, err)

	require.NoError(t, s.DeleteMovie(m.ID))
}

func TestGetAllMovies(t *testing.T) {
	s, err := json.NewStorage()
	require.NoError(t, err)

	newM1 := adding.Movie{
		Title: "New Movie 1",
		Plot:  "New movies plot 1.",
	}

	newM2 := adding.Movie{
		Title: "New Movie 2",
		Plot:  "New movies plot 2.",
	}

	m1, err := s.AddMovie(newM1)
	require.NoError(t, err)

	m2, err := s.AddMovie(newM2)
	require.NoError(t, err)

	m := s.GetAllMovies()
	require.NotEmpty(t, m)
	// TODO: use assert.EqualValues(t, m1, a) once the time bug is fixed.
	assert.ObjectsAreEqualValues(m1, m[len(m)-2])
	assert.ObjectsAreEqualValues(m2, m[len(m)-1])

	require.NoError(t, s.DeleteMovie(m1.ID))
	require.NoError(t, s.DeleteMovie(m2.ID))
}

func TestDeleteMovie(t *testing.T) {
	s, err := json.NewStorage()
	require.NoError(t, err)

	newM := adding.Movie{
		Title: "New Movie",
		Plot:  "New movies plot.",
	}

	m, err := s.AddMovie(newM)

	require.NoError(t, err)

	require.NoError(t, s.DeleteMovie(m.ID))
}
