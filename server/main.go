package main

import (
	"go.uber.org/fx"
	"go/server/config"
	"go/server/server"
)

func main() {
	fx.New(start()).Run()
}

func start() fx.Option {
	return fx.Options(config.Module, server.Module)
}
