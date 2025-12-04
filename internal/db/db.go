// Package db provides database connection and configuration.
package db

import (
	"fmt"
	"time"

	"github.com/leonardocbsr/golang-grpc-template/internal/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New creates a new database connection using the provided configuration.
func New(cfg config.Config, logger *logrus.Logger) (*gorm.DB, error) {
	c := cfg.GetDatabaseConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.Host, c.Username, c.Password, c.Database, c.Port)

	logger.Infof("Connecting to database: %s...", c.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		logger.Errorf("Failed to connect to database: %s", c.Database)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if cfg.GetLoggerConfig().Level == "debug" {
		db = db.Debug()
	}

	return db, nil
}
