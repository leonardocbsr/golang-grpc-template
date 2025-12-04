package clients

import (
	"github.com/leonardocbsr/golang-grpc-template/internal/config"
	"github.com/leonardocbsr/golang-grpc-template/internal/constants"
	ping "github.com/leonardocbsr/golang-grpc-template/pkg/ping/proto"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// NewPingClient creates a new gRPC client for the ping service.
func NewPingClient(lc fx.Lifecycle, logger *logrus.Logger, config config.Config) (ping.PingServiceClient, error) {
	return NewClient(lc, logger, config, constants.PingServiceName, ping.NewPingServiceClient)
}
