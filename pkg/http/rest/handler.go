package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
)

// Handler maps the different existing endpoints with the functions they must call
func Handler(a adding.Service, l listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/movies", listMovies(l))
	router.GET("/movies/:id", getMovie(l))
	router.POST("/movies", addMovie(a))

	return router
}
