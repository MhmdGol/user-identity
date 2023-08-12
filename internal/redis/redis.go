package redis

import (
	"Identity/cmd/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(conf config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Uri,
		Password: conf.Redis.Password,
		DB:       0,
	})
}
