package server

import (
	"context"
	"fmt"
	"github.com/PhuMinh08082001/server-cobra/config"
	"github.com/PhuMinh08082001/server-cobra/grpc/server"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
	"proto/grpc/account"
)

var Module = fx.Invoke(registerServer)

func registerServer(
	lifecycle fx.Lifecycle, configuration config.Configuration, conn net.Listener, accountServer *server.AccountServer,
) {
	grpcServer := grpc.NewServer()
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Start migration

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
				grpcServer.Stop()
				log.Println("Stop Grpc Server !")
				return nil
			},
		},
	)
}
