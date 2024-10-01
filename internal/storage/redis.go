package storage

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/zsandibe/online-course-platform/config"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	_, err := redisDb.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %v", err)
	}
	return redisDb, nil
}
