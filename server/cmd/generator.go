/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/PhuMinh08082001/server-cobra/config"
	"github.com/PhuMinh08082001/server-cobra/dal"
	"github.com/PhuMinh08082001/server-cobra/grpc/service"
	"github.com/PhuMinh08082001/server-cobra/utils/gen"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// generatorCmd represents the generator command
var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "Generate query code for database",
	Run:   generateCodegen,
}

func generateCodegen(cmd *cobra.Command, args []string) {
	fx.New(injectGen())
}

func injectGen() fx.Option {
	return fx.Options(
		config.Module,
		service.Module,
		dal.Module,
		gen.Module,
	)
}

func init() {
	rootCmd.AddCommand(generatorCmd)
}
