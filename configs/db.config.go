package configs

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DatabaseConnection string
	GormConfig *gorm.Config
	initializedDbConfig bool
)

func InitDatabaseConfig() error {
	if initializedDbConfig {
		return nil
	}

	err := InitEnvConfig()
	if err != nil {
		return err
	}

	if DatabaseConnection == "" {
		DatabaseConnection = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
			EnvConfig.DatabaseHost,
			EnvConfig.DatabaseName,
			EnvConfig.DatabaseUser,
			EnvConfig.DatabasePassword,
			EnvConfig.DatabaseName,
		)
	}

	if GormConfig == nil {
		logLevel := logger.Silent
		environment := EnvConfig.Environment
		if environment == EnvDev {
			logLevel = logger.Info
		}
	
		GormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		}
	}

	initializedDbConfig = true
	log.Printf("configs: Initialized database config\n")
	return nil
}
