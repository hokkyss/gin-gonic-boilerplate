package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkyss/gin-gonic-boilerplate/api"
	"github.com/hokkyss/gin-gonic-boilerplate/configs"
	"github.com/hokkyss/gin-gonic-boilerplate/controllers"
	"github.com/hokkyss/gin-gonic-boilerplate/models"
	"github.com/joho/godotenv"
)

func main() {
	api.GinApp.Run(":8000")
}
