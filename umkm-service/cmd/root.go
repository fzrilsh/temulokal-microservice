package cmd

import (
	"os"
	"temulokal-microservice/shared-service/logger"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "temulokal-microservice/umkm-service",
	Short: "UMKM microservice for TemuLokal",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
