package server

import (
	"cbsr.io/golang-grpc-template/common/controller"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type registerParams struct {
	fx.In

	Controllers []controller.IController `group:"controller"`
	Server      *grpc.Server
}

func RegisterControllers(p registerParams) {
	for _, c := range p.Controllers {
		c.RegisterController(p.Server)
	}
}
