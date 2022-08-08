package server

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go/server/config"
	"go/server/grpc/server"
	"google.golang.org/grpc"
	"proto/grpc/account"
)

var Module = fx.Invoke(registerServer)

func registerServer(
	lifecycle fx.Lifecycle, configuration config.Configuration,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				grpcServer := grpc.NewServer()
				accountServer := server.AccountServer{}

				account.RegisterUserServiceServer(grpcServer, &accountServer)

				fmt.Printf("Starting gRPC server at : %s:%d \n", configuration.Grpc.Host, configuration.Grpc.Port)

				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Stop Grpc Server")
				return nil
			},
		},
	)
}
