package adding

import "errors"

// ErrDuplicateMovie is used when a movie already exists.
var ErrDuplicateMovie = errors.New("movie already exists")

// Movie defines the properties of a movie to be added
type Movie struct {
	Title string `json:"title"`
	Plot  string `json:"plot"`
}

// AddMovie persists the given movie(s) to storage
func (s *service) AddMovie(m ...Movie) error {
	// make sure we don't add any duplicates
	existingMovies := s.r.GetAllMovies()
	for _, mm := range m {
		for _, e := range existingMovies {
			if mm.Title == e.Title &&
				mm.Plot == e.Plot {
				return ErrDuplicateMovie
			}
		}
	}

	// any other validation can be done here
	for _, movie := range m {
		_, _ = s.r.AddMovie(movie) // error handling omitted for simplicity
	}

	return nil
}

// AddSampleMovies adds some sample movies to the database
func (s *service) AddSampleMovies(m []Movie) {

	// any validation can be done here

	for _, mm := range m {
		_, _ = s.r.AddMovie(mm) // error handling omitted for simplicity
	}
}
