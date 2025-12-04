package main

import (
	"github.com/sirupsen/logrus"
	fxlogrus "github.com/takt-corp/fx-logrus"
	"gorm.io/gorm"

	"github.com/leonardocbsr/golang-grpc-template/internal/clients"
	"github.com/leonardocbsr/golang-grpc-template/internal/config"
	"github.com/leonardocbsr/golang-grpc-template/internal/controller"
	"github.com/leonardocbsr/golang-grpc-template/internal/db"
	"github.com/leonardocbsr/golang-grpc-template/internal/logger"
	ping "github.com/leonardocbsr/golang-grpc-template/pkg/ping/controllers"
	"github.com/leonardocbsr/golang-grpc-template/pkg/server"
	"github.com/leonardocbsr/golang-grpc-template/pkg/users/application"
	users "github.com/leonardocbsr/golang-grpc-template/pkg/users/controllers"
	"github.com/leonardocbsr/golang-grpc-template/pkg/users/models"
	"github.com/leonardocbsr/golang-grpc-template/pkg/users/repository"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"google.golang.org/grpc"
)

// runMigrations runs database migrations for all models.
func runMigrations(db *gorm.DB, logger *logrus.Logger) error {
	logger.Info("Running database migrations...")
	return db.AutoMigrate(&models.User{})
}

func main() {
	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			return &fxlogrus.LogrusLogger{Logger: logrus.New()}
		}),
		// Provide configuration
		fx.Provide(
			config.New,
			logger.New,
			db.New,
			server.New,
		),

		// Provide clients
		fx.Provide(
			clients.NewPingClient,
			clients.NewUserClient,
		),

		// Provide repositories
		fx.Provide(
			repository.NewUserRepository,
		),

		// Provide services
		fx.Provide(
			application.NewUserService,
		),

		// Provide controllers
		fx.Provide(
			controller.FxController(ping.New),
			controller.FxController(users.New),
		),
		// Register controllers
		fx.Invoke(server.RegisterControllers),

		// Run database migrations
		fx.Invoke(runMigrations),

		// Start the server
		fx.Invoke(func(*grpc.Server) {}),
	).Run()
}
