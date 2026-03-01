package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectRedis() {
	redisUrl := os.Getenv("REDIS_URL")
	if os.Getenv("ENVIRONMENT") == "development" {
		redisUrl = os.Getenv("LOCAL_REDIS_URL")
	}

	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Printf("Critical error: Unable to parse Redis URL: %v", err)
		return
	}

	Redis = redis.NewClient(opts)
	if err := Redis.Ping(context.Background()).Err(); err != nil {
		log.Printf("Redis connection error: %v", err)
		Redis = nil
	} else {
		if os.Getenv("ENVIRONMENT") == "development" {
			fmt.Println("Redis connected successfully via Localhost!")
		} else {
			fmt.Println("Redis connected successfully via Railway!")
		}
	}
}