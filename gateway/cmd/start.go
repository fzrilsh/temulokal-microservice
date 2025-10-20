package cmd

import (
	"temulokal-microservice/gateway/bootstrap"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start gateway proxy",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartProxy()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
