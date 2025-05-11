package main

import (
	"fmt"
	"github.com/alimarzban99/go-blog-api/config"
	"github.com/alimarzban99/go-blog-api/internal/middlewares"
	"github.com/alimarzban99/go-blog-api/internal/model"
	_ "github.com/alimarzban99/go-blog-api/internal/model"
	"github.com/alimarzban99/go-blog-api/internal/routers"
	"github.com/alimarzban99/go-blog-api/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	config.LoadConfig()
	err := database.InitDb()
	model.Starter()
	defer database.CloseDb()

	err = database.InitRedis()
	defer database.CloseRedis()

	gin.SetMode(config.Config.App.Env)
	router := gin.Default()

	apiV1 := router.Group("api/v1/", middlewares.Throttle())

	routers.AuthRouter(apiV1)
	routers.UserRouter(apiV1)
	routers.CategoryRouter(apiV1)
	routers.PostRouter(apiV1)

	runPort := fmt.Sprintf(":%d", config.Config.App.Port)
	err = router.Run(runPort)
	if err != nil {
		log.Fatal(err.Error())
	}
}
