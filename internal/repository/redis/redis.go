package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	redisClient *redis.Client
}

func NewRedis(redisClient *redis.Client) *Redis {
	return &Redis{redisClient: redisClient}
}

func (r *Redis) Set(key string, value string, expiration time.Duration) error {
	return nil
}

func (r *Redis) Get(key string) (string, error) {
	return "", nil
}

func (r *Redis) Delete(key string) error {
	return nil
}
