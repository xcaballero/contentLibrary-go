package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/deleting"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

// Handler maps the different existing endpoints with the functions they must call
func Handler(a adding.Service, l listing.Service, d deleting.Service) *gin.Engine {
	router := gin.Default()

	router.GET("/movies", listMovies(l))
	router.GET("/movies/:id", getMovie(l))
	router.POST("/movies", addMovie(a))
	router.DELETE("/movies/:id", deleteMovie(d))

	return router
}
