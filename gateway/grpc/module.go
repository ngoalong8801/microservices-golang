package grpc

import (
	"gateway-service/grpc/client"
	gconfig "gateway-service/grpc/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	client.Module,
	gconfig.Module,
)
