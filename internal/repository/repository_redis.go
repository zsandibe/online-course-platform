package repository

import (
	"time"

	"github.com/go-redis/redis"
	redisRepo "github.com/zsandibe/online-course-platform/internal/repository/redis"
)

type RedisRepository interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Redis struct {
	RedisRepository RedisRepository
}

func NewRedisRepository(redisClient *redis.Client) *Redis {
	return &Redis{RedisRepository: redisRepo.NewRedis(redisClient)}
}
