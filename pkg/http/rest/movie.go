package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/deleting"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

// addMovie returns a handler for POST /movies request
func addMovie(s adding.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newMovie adding.Movie
		if err := c.ShouldBindJSON(&newMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		m, err := s.AddMovie(newMovie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, m)
	}
}

// listMovies returns a handler for GET /mvoies requests
func listMovies(s listing.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, s.ListMovies())
	}
}

// getMovie returns a handler for GET /movies/:id requests
func getMovie(s listing.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		ID := c.Param("id")

		m, err := s.GetMovie(ID)
		if err == listing.ErrMovieNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, m)
	}
}

// deleteMovie returns a handler for DELETE /movies/:id requests
func deleteMovie(s deleting.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		ID := c.Param("id")

		m, err := s.DeleteMovie(ID)
		if err == deleting.ErrMovieNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, m)
	}
}
