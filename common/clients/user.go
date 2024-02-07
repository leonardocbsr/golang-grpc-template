package clients

import (
	"cbsr.io/golang-grpc-template/common/constants"
	"cbsr.io/golang-grpc-template/config"
	"cbsr.io/golang-grpc-template/proto/users"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func NewUserClient(lc fx.Lifecycle, logger *logrus.Logger, config config.IConfig) users.UserServiceClient {
	clientConfig, err := config.GetClientConfig(constants.UserServiceName)
	if err != nil {
		panic(err)
	}

	conn := openConnection(lc, logger, constants.UserServiceName, *clientConfig)
	return users.NewUserServiceClient(conn)
}
