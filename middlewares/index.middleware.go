package middlewares

import (
	"log"

	"github.com/hokkyss/gin-gonic-boilerplate/usecases"
)

var (
	initialized bool
)

func InitMiddlewares() error {
	if initialized {
		return nil
	}

	err := usecases.InitUseCases()
	if err != nil {
		return err
	}

	initialized = true
	log.Printf("middlewares: Initialized middlewares")
	return nil
}