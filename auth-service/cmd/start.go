package cmd

import (
	"temulokal-microservice/auth-service/bootstrap"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the auth service server",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
