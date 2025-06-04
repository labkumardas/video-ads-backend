package storage

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	addr := os.Getenv("REDIS_ADDR") // e.g., redis:6379

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // add if needed
		DB:       0,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}

	log.Println("✅ Connected to Redis")
}
