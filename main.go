package main

import (
	"github.com/sirupsen/logrus"
	fxlogrus "github.com/takt-corp/fx-logrus"

	"cbsr.io/golang-grpc-template/common/clients"
	"cbsr.io/golang-grpc-template/common/controller"
	"cbsr.io/golang-grpc-template/config"
	"cbsr.io/golang-grpc-template/db"
	"cbsr.io/golang-grpc-template/logger"
	ping "cbsr.io/golang-grpc-template/modules/ping/controllers"
	"cbsr.io/golang-grpc-template/modules/users/application"
	users "cbsr.io/golang-grpc-template/modules/users/controllers"
	"cbsr.io/golang-grpc-template/modules/users/repository"
	"cbsr.io/golang-grpc-template/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"google.golang.org/grpc"
)

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

		// Start the server
		fx.Invoke(func(*grpc.Server) {}),
	).Run()
}
