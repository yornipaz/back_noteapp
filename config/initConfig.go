package config

import "os"

func Init() {
	if os.Getenv("ENV") != "production" {
		loadEnvironments()
	}
	dbConfig()
	syncDatabase()
}
