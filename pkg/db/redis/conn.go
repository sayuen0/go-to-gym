package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sayuen0/go-to-gym/config"
)

// NewRedisClient return redis.Client
func NewRedisClient(cfg *config.Config) *redis.Client {
	redisHost := cfg.Redis.Addr
	if redisHost == "" {
		redisHost = ":6379"
	}

	return redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password, // no password set
		DB:           cfg.Redis.DB,       // use default DB
	})
}
