package server

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go/server/cmd/migration"
	"go/server/config"
	"go/server/grpc/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"proto/grpc/account"
)

var Module = fx.Invoke(registerServer)

func registerServer(
	lifecycle fx.Lifecycle, configuration config.Configuration, conn net.Listener, accountServer *server.AccountServer,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Start migration
				go migration.Start(configuration)

				grpcServer := grpc.NewServer()
				account.RegisterUserServiceServer(grpcServer, accountServer)

				fmt.Printf("Starting gRPC server at : %s:%d \n", configuration.Grpc.Host, configuration.Grpc.Port)

				go func() {
					if err := grpcServer.Serve(conn); err != nil {
						log.Fatal(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Stop Grpc Server")
				return nil
			},
		},
	)
}
