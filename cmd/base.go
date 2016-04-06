package cmd

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/yiziz/gin-template/app/models"
	"github.com/yiziz/gin-template/config"
	"github.com/yiziz/gin-template/config/initializers"
)

func initDB() *gorm.DB {
	config.SetAppEnv()
	db := initializers.ConnectDB(config.AppEnv)
	models.SetDB(db)
	return db
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "run server",
	Long:  `run server`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Run()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
