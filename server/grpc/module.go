package grpc

import (
	"github.com/PhuMinh08082001/server-cobra/grpc/repository"
	"github.com/PhuMinh08082001/server-cobra/grpc/server"
	"github.com/PhuMinh08082001/server-cobra/grpc/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	repository.Module,
	service.Module,
	server.Module,
)
