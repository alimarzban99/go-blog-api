package commands

import (
	"fmt"
	"github.com/alimarzban99/go-blog-api/config"
	"github.com/alimarzban99/go-blog-api/internal/middlewares"
	"github.com/alimarzban99/go-blog-api/internal/model"
	"github.com/alimarzban99/go-blog-api/internal/routers"
	"github.com/alimarzban99/go-blog-api/pkg/database"
	"github.com/alimarzban99/go-blog-api/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
)

var port int

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {

		config.LoadConfig()
		logger := logging.NewLogger()
		err := database.InitDb()
		if err != nil {
			logger.Fatal(logging.Startup, err.Error())
		}
		defer database.CloseDb()

		err = database.InitRedis()
		if err != nil {
			logger.Fatal(logging.Startup, err.Error())
		}
		defer database.CloseRedis()

		model.Starter()

		gin.SetMode(config.Config.App.Env)
		router := gin.Default()

		apiV1 := router.Group("api/v1/", middlewares.Throttle())
		routers.AuthRouter(apiV1)
		routers.UserRouter(apiV1)
		routers.CategoryRouter(apiV1)
		routers.PostRouter(apiV1)

		runPort := fmt.Sprintf(":%d", config.Config.App.Port)
		if port != 0 {
			runPort = fmt.Sprintf(":%d", port)
		}
		log.Printf("Server is running at http://localhost%s\n", runPort)
		err = router.Run(runPort)
		if err != nil {
			log.Fatal(err)
		}
	},
}

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

func init() {
	ServeCmd.Flags().IntVarP(&port, "port", "p", 0, "Port to run the server on")
}
