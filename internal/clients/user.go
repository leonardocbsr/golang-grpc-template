package clients

import (
	"github.com/leonardocbsr/golang-grpc-template/internal/config"
	"github.com/leonardocbsr/golang-grpc-template/internal/constants"
	users "github.com/leonardocbsr/golang-grpc-template/pkg/users/proto"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// NewUserClient creates a new gRPC client for the user service.
func NewUserClient(lc fx.Lifecycle, logger *logrus.Logger, config config.Config) (users.UserServiceClient, error) {
	return NewClient(lc, logger, config, constants.UserServiceName, users.NewUserServiceClient)
}
