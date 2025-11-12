package cmd

import (
	"temulokal-microservice/umkm-service/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start UMKM service server",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartServer()
	},
}
