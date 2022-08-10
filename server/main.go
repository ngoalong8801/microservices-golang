package main

import (
	"go.uber.org/fx"
	"go/server/cmd/migration"
	"go/server/config"
	"go/server/dal"
	"go/server/grpc"
	"go/server/server"
)

func main() {
	fx.New(start()).Run()
}

func start() fx.Option {
	return fx.Options(
		dal.Module,
		grpc.Module,
		config.Module,
		migration.Module,
		server.Module,
	)
}
