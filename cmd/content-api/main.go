package main

import (
	"fmt"
	"log"

	"github.com/xcaballero/contentLibrary-go/pkg/adding"
	"github.com/xcaballero/contentLibrary-go/pkg/deleting"
	"github.com/xcaballero/contentLibrary-go/pkg/http/rest"
	"github.com/xcaballero/contentLibrary-go/pkg/listing"
	"github.com/xcaballero/contentLibrary-go/pkg/repository"
	"github.com/xcaballero/contentLibrary-go/pkg/storage/json"
	"github.com/xcaballero/contentLibrary-go/pkg/storage/memory"
)

// StorageType defines available storage types
type StorageType int

const (
	// JSON will store data in JSON files saved on disk
	JSON StorageType = iota
	// Memory will sotre data in memory
	Memory
)

func main() {
	// Set up storage
	storageType := JSON

	var repo repository.Repository
	var adder adding.Service
	var lister listing.Service
	var deleter deleting.Service

	switch storageType {
	case Memory:
		storage := new(memory.Storage)
		cache := new(memory.Storage)

		repo = repository.NewRepository(storage, cache)
		adder = adding.NewService(repo)
		lister = listing.NewService(repo)
		deleter = deleting.NewService(repo)

	case JSON:
		// error handling omitted for simplicity.
		storage, _ := json.NewStorage()
		cache, _ := json.NewStorage()

		repo = repository.NewRepository(storage, cache)
		adder = adding.NewService(repo)
		lister = listing.NewService(repo)
		deleter = deleting.NewService(repo)
	}

	// set up the HTTP server
	router := rest.Handler(adder, lister, deleter)

	fmt.Println("The content library server is on tap now: https://localhost:5000")
	log.Fatal(router.Run(":5000"))
}
