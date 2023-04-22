package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkyss/gin-gonic-boilerplate/configs"
	"github.com/hokkyss/gin-gonic-boilerplate/controllers"
	"github.com/hokkyss/gin-gonic-boilerplate/models"
	"github.com/joho/godotenv"
)

var (
	ginApp *gin.Engine
)

func init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	err = configs.InitEnvConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// err = configs.InitDatabaseConfig()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err = models.InitModels()
	if err != nil {
		log.Fatalln(err)
	}

	if configs.EnvConfig.Environment == configs.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}
	if configs.EnvConfig.Environment == configs.EnvTest {
		gin.SetMode(gin.TestMode)
	}

	ginApp = gin.New()
	v1Api := ginApp.Group("/api/v1")

	err = controllers.InitControllers(v1Api)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	ginApp.Run(":8000")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ginApp.ServeHTTP(w, r)
}
