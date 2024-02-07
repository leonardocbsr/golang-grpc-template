package clients

import (
	"context"

	"cbsr.io/golang-grpc-template/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func openConnection(lc fx.Lifecycle, logger *logrus.Logger, name string, clientConfig config.ClientsConfig) *grpc.ClientConn {
	logger.Infof("Opening connection to %s gRPC server...", name)
	conn, err := grpc.Dial(clientConfig.GetURL(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Errorf("Failed to open connection to %s gRPC server: %v", name, err)
		panic(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Infof("Closing connection to %s gRPC server...", name)
			conn.Close()
			return nil
		},
	})

	return conn
}
