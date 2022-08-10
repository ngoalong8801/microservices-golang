package grpc

import (
	"go.uber.org/fx"
	"go/server/grpc/repository"
	"go/server/grpc/server"
	"go/server/grpc/service"
)

var Module = fx.Options(
	repository.Module,
	service.Module,
	server.Module,
)
