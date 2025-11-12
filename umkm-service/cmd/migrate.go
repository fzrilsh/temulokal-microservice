package cmd

import (
	"temulokal-microservice/umkm-service/config"
	"temulokal-microservice/umkm-service/utils/database"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations for UMKM service",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		database.Migrate(cfg)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
