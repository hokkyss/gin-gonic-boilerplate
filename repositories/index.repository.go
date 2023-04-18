package repositories

import "log"

var (
	initialized bool
)

func InitRepository() error {
	if initialized {
		return nil
	}

	initialized = true
	log.Printf("repositories: Initialized repositories")
	return nil
}