package clients

import (
	"cbsr.io/golang-grpc-template/common/constants"
	"cbsr.io/golang-grpc-template/config"
	"cbsr.io/golang-grpc-template/proto/ping"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func NewPingClient(lc fx.Lifecycle, logger *logrus.Logger, config config.IConfig) ping.PingServiceClient {
	clientConfig, err := config.GetClientConfig(constants.PingServiceName)
	if err != nil {
		panic(err)
	}

	conn := openConnection(lc, logger, constants.PingServiceName, *clientConfig)
	return ping.NewPingServiceClient(conn)
}
