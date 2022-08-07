package main

import (
	"gateway-service/config"
	"gateway-service/controller"
	"gateway-service/grpc"
	"gateway-service/routes"
	"gateway-service/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(start()).Run()
}

func start() fx.Option {
	return fx.Options(
		config.Module,
		grpc.Module,
		controller.Module,
		routes.Module,
		server.Module,
	)
}
