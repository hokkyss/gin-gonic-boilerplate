package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hokkyss/gin-gonic-boilerplate/types"
)

func healthCheckRoutes(router *gin.RouterGroup)  {
	router.GET("/ping", handlePing)
}

func handlePing(c *gin.Context) {
	c.Status(200)
	c.JSON(200, types.PingResponse{Pong: "Pong"})
}
