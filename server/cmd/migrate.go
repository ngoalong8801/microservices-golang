/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/PhuMinh08082001/server-cobra/config"
	"github.com/PhuMinh08082001/server-cobra/dal"
	"github.com/PhuMinh08082001/server-cobra/utils/migrate"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// migrateCmd represents the migration command
var migrateCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Run:   initMigrate,
}

func initMigrate(cmd *cobra.Command, args []string) {
	fx.New(injectMigrate())
}

func injectMigrate() fx.Option {
	return fx.Options(
		dal.Module,
		config.Module,
		migrate.Module,
	)

}
func init() {
	rootCmd.AddCommand(migrateCmd)

}
