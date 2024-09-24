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
	cfg, err := config.NewConfig(".env")
	if err != nil {
		return fmt.Errorf("config.NewConfig: %v", err)
	}

	db, err := storage.NewPostgresDB(cfg)
	if err != nil {
		return fmt.Errorf("storage.NewPostgresDB: %v", err)
	}
	defer db.Close()
	logger.Info("Database  loaded successfully")

	if err = db.MigrateUp(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Debug(err)
	}

	hash := hash.NewHash()

	tokenManager, err := manager.NewManager(cfg.Token.SigningKey)
	if err != nil {
		return fmt.Errorf("manager.NewManager: %v", err)
	}

	repo := repository.NewPostgresRepository(db.DB)

	svc := service.NewService(repo, tokenManager, hash, cfg.Token.AccessTokenTTL, cfg.Token.RefreshTokenTTL)

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
