package main

import (
	"log"

	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/deleting"
	"github.com/xcaballero/contentLibrary-go/pkg/http/rest"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
	"github.com/xcaballero/contentLibrary-go/pkg/repository"
	"github.com/xcaballero/contentLibrary-go/pkg/storage/json"
	"github.com/xcaballero/contentLibrary-go/pkg/storage/redis"
)

func main() {
	storage, err := redis.NewStorage()
	if err != nil {
		panic(err)
	}
	cache, err := json.NewStorage()
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(storage, cache)

	adder := adding.NewService(repo)
	lister := listing.NewService(repo)
	deleter := deleting.NewService(repo)

	// set up the HTTP server
	router := rest.Handler(adder, lister, deleter)
	log.Fatal(router.Run(":5000"))
}
