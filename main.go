package main

import (
	"github.com/hokkyss/gin-gonic-boilerplate/api"
)

func main() {
	api.Init()
	api.GinApp.Run(":8000")
}
