package configs

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

type EnvironmentType = string
const (
	EnvDev EnvironmentType = "development"
	EnvProd EnvironmentType = "production"
	EnvTest EnvironmentType = "test"
)

type EnvConfigType struct {
	DatabaseHost string
	DatabasePort string
	DatabaseUser string
	DatabasePassword string
	DatabaseName string
	Environment EnvironmentType
}

var (
	EnvConfig EnvConfigType
	initializedEnvConfig bool
)

func InitEnvConfig () error {
	if initializedEnvConfig {
		return nil
	}

	EnvConfig = EnvConfigType{
		DatabaseHost: os.Getenv("DATABASE_HOST"),
		DatabasePort: os.Getenv("DATABASE_PORT"),
		DatabaseUser: os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		Environment: os.Getenv("ENVIRON"),
	}

	if EnvConfig.Environment != EnvDev && EnvConfig.Environment != EnvProd && EnvConfig.Environment != EnvTest {
		EnvConfig.Environment = EnvDev
	}
	
	values := reflect.ValueOf(EnvConfig)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).String() == "" {
			fmt.Printf("configs: %s environment variables cannot be empty\n", types.Field(i).Name)
			return fmt.Errorf("configs: Initialization failed due to missing environment variables")
		}
	}

	initializedEnvConfig = true
	log.Printf("configs: Initialized environment config\n")
	return nil
}
