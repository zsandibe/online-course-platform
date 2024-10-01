package service

import (
	"time"

	"github.com/zsandibe/online-course-platform/internal/repository"
	"github.com/zsandibe/online-course-platform/internal/service/svc"
	"github.com/zsandibe/online-course-platform/pkg/hash"
	"github.com/zsandibe/online-course-platform/pkg/manager"
)

type UserSvc interface{}

type Service struct {
	UserSvc UserSvc
}

func NewService(postgresRepo *repository.Repository, redisRepo *repository.Redis, tokenManager *manager.Manager, hash *hash.PasswordHasher, accessTokenTTL, refreshTokenTTL time.Duration) *Service {
	return &Service{
		UserSvc: svc.NewUserService(tokenManager, hash, redisRepo.RedisRepository, postgresRepo.UserRepository, accessTokenTTL, refreshTokenTTL),
	}
}
