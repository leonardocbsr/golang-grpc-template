package controller

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type IController interface {
	// RegisterController registers the services with the gRPC server
	RegisterController(s *grpc.Server)
}

func FxController(f any) any {
	return fx.Annotate(f, fx.As(new(IController)), fx.ResultTags(`group:"controller"`))
}
