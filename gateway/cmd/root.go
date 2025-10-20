package cmd

import (
	"os"

	"temulokal-microservice/shared-service/logger"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "temulokal-microservice/gateway",
	Short: "Gateway proxy for TemuLokal Microservice",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
