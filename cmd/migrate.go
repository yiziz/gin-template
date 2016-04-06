package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yiziz/gin-template/db/migrate"
)

func init() {
	RootCmd.AddCommand(dbCmd)
}

var dbCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "migrate database",
	Long:  `migrate database`,
	Run: func(cmd *cobra.Command, args []string) {
		db := initDB()
		migrate.RunMigrations(db)
	},
}
