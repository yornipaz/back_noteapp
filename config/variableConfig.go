package config

import (
	"log"

	"github.com/joho/godotenv"
)

func loadVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err.Error())
	}

}
