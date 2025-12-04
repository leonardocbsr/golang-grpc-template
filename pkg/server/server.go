// Package server provides gRPC server initialization and lifecycle management.
package server

import (
	"context"
	"net"
	"time"

	"github.com/leonardocbsr/golang-grpc-template/internal/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// New creates a new gRPC server with the configured options and lifecycle hooks.
func New(lc fx.Lifecycle, log *logrus.Logger, config config.Config) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(1000),
		grpc.ConnectionTimeout(30*time.Second),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     5 * time.Minute,
			MaxConnectionAge:      30 * time.Minute,
			MaxConnectionAgeGrace: 5 * time.Second,
			Time:                  2 * time.Minute,
			Timeout:               20 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             1 * time.Minute,
			PermitWithoutStream: true,
		}),
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", config.GetServerURL())
			if err != nil {
				return err
			}

			log.Info("Registering reflection service...")
			reflection.Register(grpcServer)

			log.Infof("Starting gRPC server at port %d...", config.GetServerConfig().Port)
			go func() {
				if err := grpcServer.Serve(ln); err != nil {
					log.Errorf("gRPC server error: %v", err)
				}
			}()

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
