package svc

import (
	"time"

	"github.com/zsandibe/online-course-platform/internal/repository"
	"github.com/zsandibe/online-course-platform/pkg/hash"
	"github.com/zsandibe/online-course-platform/pkg/manager"
)

type UserService struct {
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	tokenManager    *manager.Manager
	hash            *hash.PasswordHasher
	redisRepo       repository.RedisRepository
	userRepo        repository.UserRepository
}

func NewUserService(tokenManager *manager.Manager, hash *hash.PasswordHasher, redisRepo repository.RedisRepository, userRepo repository.UserRepository, accessTokenTTL, refreshTokenTTL time.Duration) *UserService {
	return &UserService{
		tokenManager:    tokenManager,
		hash:            hash,
		redisRepo:       redisRepo,
		userRepo:        userRepo,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}
