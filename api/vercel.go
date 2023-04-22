package api

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
	GinApp *gin.Engine
)

func Init() {
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

	GinApp = gin.New()
	v1Api := GinApp.Group("/api/v1")

	err = controllers.InitControllers(v1Api)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	Init()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	GinApp.ServeHTTP(w, r)
}
