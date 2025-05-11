package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alimarzban99/go-blog-api/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var redisClient *redis.Client

func InitRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Config.Redis.Host, config.Config.Redis.Port),
		Password:     config.Config.Redis.Password,
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolSize:     10,
		PoolTimeout:  15 * time.Second,
	})

	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	err := redisClient.Close()
	if err != nil {
		log.Println(err.Error())
	}
}

func Set[T any](ctx context.Context, c *redis.Client, key string, value T, duration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, v, duration).Err()
}

func Get[T any](ctx context.Context, c *redis.Client, key string) (T, error) {
	var dest T = *new(T)
	v, err := c.Get(ctx, key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(v), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}
