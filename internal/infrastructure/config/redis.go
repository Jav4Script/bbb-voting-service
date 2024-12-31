package config

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("REDIS_URL is not set")
	}

	options, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	options.DialTimeout = 30 * time.Second
	options.ReadTimeout = 30 * time.Second
	options.WriteTimeout = 30 * time.Second

	client := redis.NewClient(options)

	return client
}
