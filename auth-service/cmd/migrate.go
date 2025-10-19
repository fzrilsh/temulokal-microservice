package cmd

import (
	"github.com/spf13/cobra"
	"temulokal-microservice/auth-service/config"
	"temulokal-microservice/auth-service/utils/database"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		database.Migrate(cfg)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
