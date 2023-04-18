package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hokkyss/gin-gonic-boilerplate/models"
)

func healthCheckRoutes(router *gin.RouterGroup)  {
	router.GET("/ping", handlePing)
}

func handlePing(c *gin.Context) {
	c.Status(200)
	c.JSON(200, models.PingResponse{Pong: "Pong"})
}
