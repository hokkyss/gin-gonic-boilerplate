package models

import (
	"fmt"
	"log"

	"github.com/hokkyss/gin-gonic-boilerplate/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	initialized bool
)

func GetDatabase() (*gorm.DB, error) {
	if database != nil {
		return database, nil
	}

	db, err := gorm.Open(postgres.Open(configs.DatabaseConnection), configs.GormConfig)
	if err != nil {
		return nil, fmt.Errorf("models: Failed connecting to database")
	}

	database = db
	return db, err
}

func InitModels() error {
	if initialized {
		return nil
	}

	// db, err := GetDatabase()
	// if err != nil {
	// 	return err
	// }

	// err = db.AutoMigrate(&User{})
	// if err != nil {
	// 	return fmt.Errorf("models: Migration failed")
	// }

	initialized = true
	log.Printf("models: Initialized models\n")
	return nil
}
