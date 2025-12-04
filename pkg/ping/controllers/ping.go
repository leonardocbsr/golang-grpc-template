// Package controller implements the PingService gRPC controller.
//
//go:generate protoc --go_out=../proto --go_opt=paths=source_relative --go-grpc_out=../proto --go-grpc_opt=paths=source_relative -I=../proto ping.proto
package controller

import (
	"context"

	"github.com/leonardocbsr/golang-grpc-template/internal/controller"
	ping "github.com/leonardocbsr/golang-grpc-template/pkg/ping/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var _ ping.PingServiceServer = (*pingServer)(nil)
var _ controller.Controller = (*pingServer)(nil)

type pingServer struct {
	ping.UnimplementedPingServiceServer
	logger *logrus.Entry
}

func New(l *logrus.Logger) controller.Controller {
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
