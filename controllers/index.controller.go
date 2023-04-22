package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hokkyss/gin-gonic-boilerplate/middlewares"
	"github.com/hokkyss/gin-gonic-boilerplate/usecases"
)

var (
	initialized bool
)

func InitControllers (router *gin.RouterGroup) error {
	if initialized {
		return nil
	}

	err := usecases.InitUseCases()
	if err != nil {
		return err
	}

	err = middlewares.InitMiddlewares()
	if err != nil {
		return err
	}

	healthCheckRoutes(router)

	initialized = true
	log.Printf("controllers: Initialized route handlers")
	return nil
}
