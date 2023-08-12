package redis

import (
	"Identity/cmd/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(conf config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Uri,
		Password: conf.Redis.Password,
		DB:       0,
	})
}
