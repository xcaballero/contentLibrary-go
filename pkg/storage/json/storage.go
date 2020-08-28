package json

import (
	"path"
	"runtime"

	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	// dir defines the name of the directory where the files are stored.
	dir = "/data/"

	// CollectionMovie identifies for the JSON collection of movies
	CollectionMovie = "movies"
)

// Storage stores data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)

	return s, err
}
