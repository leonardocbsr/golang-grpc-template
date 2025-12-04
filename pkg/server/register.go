package server

import (
	"github.com/leonardocbsr/golang-grpc-template/internal/controller"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type registerParams struct {
	fx.In

	Controllers []controller.Controller `group:"controller"`
	Server      *grpc.Server
}

// RegisterControllers registers all controllers with the gRPC server.
func RegisterControllers(p registerParams) {
	for _, c := range p.Controllers {
		c.RegisterController(p.Server)
	}
}
