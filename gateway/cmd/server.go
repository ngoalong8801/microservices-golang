/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gateway-service/config"
	"gateway-service/controller"
	"gateway-service/grpc"
	"gateway-service/routes"
	"gateway-service/server"
	"go.uber.org/fx"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Run:   initServer,
}

func initServer(cmd *cobra.Command, args []string) {
	fx.New(inject()).Run()
}

func inject() fx.Option {
	return fx.Options(
		config.Module,
		grpc.Module,
		controller.Module,
		routes.Module,
		server.Module,
	)
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
