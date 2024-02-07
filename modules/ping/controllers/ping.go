package controller

import (
	"context"

	"cbsr.io/golang-grpc-template/common/controller"
	"cbsr.io/golang-grpc-template/proto/ping"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var _ ping.PingServiceServer = (*pingServer)(nil)
var _ controller.IController = (*pingServer)(nil)

type pingServer struct {
	ping.UnimplementedPingServiceServer
	logger *logrus.Entry
}

func New(l *logrus.Logger) controller.IController {
	return &pingServer{
		logger: l.WithField("service", "ping"),
	}
}

func (s *pingServer) RegisterController(srv *grpc.Server) {
	ping.RegisterPingServiceServer(srv, s)
}

func (s *pingServer) Ping(ctx context.Context, req *ping.PingRequest) (*ping.PongResponse, error) {
	return &ping.PongResponse{Reply: "pong"}, nil
}
