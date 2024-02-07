package server

import (
	"context"
	"net"

	"cbsr.io/golang-grpc-template/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(lc fx.Lifecycle, log *logrus.Logger, config config.IConfig) *grpc.Server {
	grpcServer := grpc.NewServer()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", config.GetServerURL())
			if err != nil {
				return err
			}

			log.Info("Registering reflection service...")
			reflection.Register(grpcServer)

			log.Infof("Starting gRPC server at port %d...", config.GetServerConfig().Port)
			go grpcServer.Serve(ln)

			log.Info("Listening for gRPC connections")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping gRPC server...")
			grpcServer.GracefulStop()
			return nil
		},
	})
	return grpcServer
}
