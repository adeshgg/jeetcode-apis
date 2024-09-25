package cache

import (
	"context"
	"jeetcode-apis/config"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Connect(cfg *config.Config) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		// log.Fatalf("Could not connect to Redis: %v", err)
		return &redis.Client{}, err
	}

	return rdb, nil
}
