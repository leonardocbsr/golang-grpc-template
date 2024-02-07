package db

import (
	"fmt"

	"cbsr.io/golang-grpc-template/config"
	"cbsr.io/golang-grpc-template/modules/users/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(config config.IConfig, logger *logrus.Logger) *gorm.DB {
	c := config.GetDatabaseConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.Host, c.Username, c.Password, c.Database, c.Port)

	logger.Infof("Connecting to database: %s...", c.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to database: %s", c.Database)
		panic(err)
	}

	logger.Infof("Auto migrating database...")
	db.AutoMigrate(&models.User{})

	if config.GetLoggerConfig().Level == "debug" {
		db = db.Debug()
	}

	return db
}
