package main

import (
	"context"
	"time"

	"github.com/senn404/bookmark-managent/internal/pkg/redis"
)

func main() {
	redisClient, err := redis.NewClient("")
	if err != nil {
		panic(err)
	}

	redisClient.Set(context.Background(), "test", "test_value", time.Hour)
}
