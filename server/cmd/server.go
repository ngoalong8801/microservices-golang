/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/PhuMinh08082001/server-cobra/config"
	"github.com/PhuMinh08082001/server-cobra/dal"
	"github.com/PhuMinh08082001/server-cobra/grpc"
	"github.com/PhuMinh08082001/server-cobra/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Run:   initGrpcServer,
}

func initGrpcServer(cmd *cobra.Command, args []string) {
	fx.New(inject()).Run()
}

func inject() fx.Option {
	return fx.Options(
		dal.Module,
		grpc.Module,
		config.Module,
		server.Module,
	)
}
func init() {
	rootCmd.AddCommand(serverCmd)
}
