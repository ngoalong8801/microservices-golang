/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/PhuMinh08082001/server-cobra/config"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server-cobra",
	Short: "go-starter Management CLI",
	Run: func(cmd *cobra.Command, args []string) {
		generatorCmd.Run(cmd, args)
		serverCmd.Run(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
