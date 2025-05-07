package database

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "secret",
	DB:       0,
})
