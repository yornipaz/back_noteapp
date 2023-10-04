package config

import (
	"os"

	"github.com/yornifpaz/back_noteapp/app/models"
)

func getDefaultDatabaseConfig() (config models.DatabaseConfig) {

	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	user := string(os.Getenv("DB_USERNAME"))
	password := os.Getenv("DB_PASSWORD")
	config = models.DatabaseConfig{
		Host:     host,
		Password: password,
		Port:     port,
		User:     user,
		DBName:   database,
	}

	return

}
