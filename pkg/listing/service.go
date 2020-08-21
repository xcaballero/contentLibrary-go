package listing

// Repository provides access to the storage.
type Repository interface {
	// GetMovie returns the movie with given ID.
	GetMovie(string) (Movie, error)
	// GetAllMovies returns all movies saved in storage.
	GetAllMovies() []Movie
}

// Service provides listing operations.
type Service interface {
	// GetMovie returns the movie with given ID.
	GetMovie(string) (Movie, error)
	// GetAllMovies returns all movies saved in storage.
	GetMovies() []Movie
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
