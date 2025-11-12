package cmd

import (
	"fmt"
	"os"
	"temulokal-microservice/umkm-service/bootstrap"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umkm-service",
	Short: "UMKM service CLI",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
