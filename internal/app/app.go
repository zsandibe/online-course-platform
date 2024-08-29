package app

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/zsandibe/online-course-platform/config"
	"github.com/zsandibe/online-course-platform/internal/storage"
	logger "github.com/zsandibe/online-course-platform/pkg"
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

	fmt.Println(db)

	return nil
}
