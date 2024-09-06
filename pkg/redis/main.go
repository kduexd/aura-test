package redis

import (
	"aura-test/pkg/config"
	"aura-test/pkg/log"
	"context"

	"github.com/redis/go-redis/v9"
)

const (
	UserDB  = 0
	CacheDB = 1
)

// NewClient returns a client to the Redis Server
func NewClient(ctx context.Context, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.EnvForge().GetString("redis.PORT.CACHE.HOST"),
		Password: config.EnvForge().GetString("redis.PORT.CACHE.PASSWORD"),
		DB:       db,
		PoolSize: 2000,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Error(err)
		return nil
	}

	return client
}
