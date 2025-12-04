// Package controller provides interfaces for gRPC controller registration.
package controller

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// Controller defines the interface for gRPC service controllers.
type Controller interface {
	// RegisterController registers the services with the gRPC server.
	RegisterController(s *grpc.Server)
}

// FxController wraps a controller constructor for fx dependency injection.
func FxController(f any) any {
	return fx.Annotate(f, fx.As(new(Controller)), fx.ResultTags(`group:"controller"`))
}
