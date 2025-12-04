package clients

import (
	"context"
	"time"

	"github.com/leonardocbsr/golang-grpc-template/internal/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func openConnection(lc fx.Lifecycle, logger *logrus.Logger, name string, clientConfig config.ClientsConfig) (*grpc.ClientConn, error) {
	logger.Infof("Opening connection to %s gRPC server...", name)
	conn, err := grpc.NewClient(
		clientConfig.GetURL(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                2 * time.Minute,
			Timeout:             20 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(10*1024*1024),
			grpc.MaxCallSendMsgSize(10*1024*1024),
		),
	)
	if err != nil {
		logger.Errorf("Failed to open connection to %s gRPC server: %v", name, err)
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Infof("Closing connection to %s gRPC server...", name)
			conn.Close()
			return nil
		},
	})

	return conn, nil
}

func NewClient[T any](lc fx.Lifecycle, logger *logrus.Logger, config config.Config, serviceName string, clientFn func(grpc.ClientConnInterface) T) (T, error) {
	clientConfig, err := config.GetClientConfig(serviceName)
	if err != nil {
		var zero T
		return zero, err
	}

	conn, err := openConnection(lc, logger, serviceName, *clientConfig)
	if err != nil {
		var zero T
		return zero, err
	}

	return clientFn(conn), nil
}
