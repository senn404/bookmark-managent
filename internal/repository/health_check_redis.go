package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type HealthCheckRedis interface {
	HealthCheck(ctx context.Context) error
}

type healthCheckRedis struct {
	redisClient *redis.Client
}

func NewHealthCheckRedis(redisClient *redis.Client) HealthCheckRedis {
	return &healthCheckRedis{redisClient: redisClient}
}

//go:generate mockery --name HealthCheckRedis --filename HealthCheckRedis.go
func (h *healthCheckRedis) HealthCheck(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return h.redisClient.Ping(ctx).Err()
}
