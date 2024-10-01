package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/zsandibe/online-course-platform/config"
	v1 "github.com/zsandibe/online-course-platform/internal/delivery/api/v1"
	"github.com/zsandibe/online-course-platform/internal/delivery/server"
	"github.com/zsandibe/online-course-platform/internal/repository"
	"github.com/zsandibe/online-course-platform/internal/service"
	"github.com/zsandibe/online-course-platform/internal/storage"
	logger "github.com/zsandibe/online-course-platform/pkg"
	"github.com/zsandibe/online-course-platform/pkg/hash"
	"github.com/zsandibe/online-course-platform/pkg/manager"
)

func Start() error {
	ctx := context.Background()

	cfg, err := config.NewConfig(".env")
	if err != nil {
		return fmt.Errorf("config.NewConfig: %v", err)
	}

	redisClient, err := storage.NewRedisClient(cfg)
	if err != nil {
		return fmt.Errorf("storage.NewRedisClient: %v", err)
	}
	logger.Info("Redis client loaded successfully")

	db, err := storage.NewPostgresDB(cfg)
	if err != nil {
		return fmt.Errorf("storage.NewPostgresDB: %v", err)
	}
	defer db.Close()
	logger.Info("Database  loaded successfully")

	if err = db.MigrateUp(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Debug(err)
	}
	logger.Info("Migrations completed successfully")

	s3, err := storage.NewS3Client(ctx, cfg)
	if err != nil {
		return fmt.Errorf("storage.NewS3Client: %v", err)
	}
	logger.Info("S3 client loaded successfully")
	fmt.Println(s3)

	hash := hash.NewHash()

	tokenManager, err := manager.NewManager(cfg.Token.SigningKey)
	if err != nil {
		return fmt.Errorf("manager.NewManager: %v", err)
	}

	postgresRepo := repository.NewPostgresRepository(db.DB)

	redisRepo := repository.NewRedisRepository(redisClient)

	svc := service.NewService(postgresRepo, redisRepo, tokenManager, hash, cfg.Token.AccessTokenTTL, cfg.Token.RefreshTokenTTL)

	delivery := v1.NewHandler(svc, tokenManager)

	server := server.NewServer(cfg, delivery.Routes())
	go func() {
		if err := server.Run(); err != nil {
			logger.Error("failed to start server: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	select {
	case <-quit:
		logger.Info("Received interrupt signal. Shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("Error during server shutdown: ", err)
		}

		logger.Info("Server gracefully stopped")
	}

	fmt.Println(db)

	return nil
}
