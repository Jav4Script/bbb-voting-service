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

	// Redis adjustments
	options.DialTimeout = 15 * time.Second
	options.ReadTimeout = 15 * time.Second
	options.WriteTimeout = 15 * time.Second
	options.PoolSize = 50     // Connection pool size
	options.MinIdleConns = 20 // Minimum number of idle connections which is useful when establishing new connection is slow

	client := redis.NewClient(options)

	return client
}
