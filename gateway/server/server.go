package server

import (
	"context"
	"gateway-service/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Invoke(registerServer)

func registerServer(
	lifecycle fx.Lifecycle, configuration config.Configuration, route *gin.Engine,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go route.Run(":" + configuration.Port)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}
