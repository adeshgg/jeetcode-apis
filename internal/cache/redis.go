package cache

import (
	"context"
	"fmt"
	"jeetcode-apis/config"
	"log"

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

func EnqueueTask(rdb *redis.Client, key string, task string) error {
	if err := rdb.LPush(ctx, key, task).Err(); err != nil {
		return fmt.Errorf("could not add task to queue: %v", err)
	}

	log.Printf("Task: %s added to the queue", task)
	return nil
}

func DequeueTask(rdb *redis.Client, key string) (string, error) {
	task, err := rdb.RPop(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("could not pop task from the queue: %s, error: %v", key, err)
	}

	return task, nil
}
