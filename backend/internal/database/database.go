// Package database opens and prepares the PostgreSQL connection.
package database

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const retryInterval = 2 * time.Second

// Connect opens a connection pool, retrying until the database accepts
// connections or cfg.ConnectTimeout elapses. Retrying covers the window at
// boot where the backend starts before the database container is ready.
func Connect(ctx context.Context, cfg config.Database, log *slog.Logger) (*gorm.DB, error) {
	ctx, cancel := context.WithTimeout(ctx, cfg.ConnectTimeout)
	defer cancel()

	gormConfig := &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)}
	for attempt := 1; ; attempt++ {
		db, err := gorm.Open(postgres.Open(dsn(cfg)), gormConfig)
		if err == nil {
			if err = configurePool(db); err == nil {
				return db, nil
			}
		}

		log.Warn("database not ready", "attempt", attempt, "error", err)
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("connect to database: %w", err)
		case <-time.After(retryInterval):
		}
	}
}

// Migrate creates or updates the schema for the given models.
func Migrate(db *gorm.DB, models ...any) error {
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}
	return nil
}

func dsn(cfg config.Database) string {
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     net.JoinHostPort(cfg.Host, cfg.Port),
		Path:     cfg.Name,
		RawQuery: url.Values{"sslmode": {cfg.SSLMode}}.Encode(),
	}
	return u.String()
}

func configurePool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	return nil
}
