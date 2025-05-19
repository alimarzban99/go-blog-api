package commands

import (
	"github.com/alimarzban99/go-blog-api/config"
	"github.com/alimarzban99/go-blog-api/internal/model"
	"github.com/alimarzban99/go-blog-api/pkg/database"
	"github.com/alimarzban99/go-blog-api/pkg/logging"
	"github.com/spf13/cobra"
	"log"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Database",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
		logger := logging.NewLogger()

		err := database.InitDb()
		if err != nil {
			logger.Fatal(logging.Migration, err.Error())
		}
		defer database.CloseDb()
		model.Starter()
		log.Println("migrate database successfully")
	},
}
