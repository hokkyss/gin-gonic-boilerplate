package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkyss/gin-gonic-boilerplate/middlewares"
	"github.com/hokkyss/gin-gonic-boilerplate/types"
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
	router.Use(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, types.MessageResponse{Detail: "Not Found!"})
	})

	initialized = true
	log.Printf("controllers: Initialized route handlers")
	return nil
}
