package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const (
	// Addr defines the address to connect to the redis instance
	Addr = "localhost:6379"
	// Password defines the password to connect to the redis instance
	Password = ""
	// DB defines the database in the redis instance.
	DB = 1
)

// Storage stores data in Redis db
type Storage struct {
	db  *redis.Client
	ctx context.Context
}

// NewStorage returns a new Redis storage
func NewStorage() (*Storage, error) {
	s := new(Storage)

	s.db = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})

	s.ctx = context.Background()

	_, err := s.db.Ping(s.ctx).Result()

	return s, err
}
