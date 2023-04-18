package usecases

import (
	"log"

	"github.com/hokkyss/gin-gonic-boilerplate/repositories"
)

var (
	initialized bool
)

func InitUseCases() error {
	if initialized {
		return nil
	}

	err := repositories.InitRepository()
	if err != nil {
		return err
	}

	initialized = true
	log.Printf("usecases: Initialized usecases")
	return nil
}
