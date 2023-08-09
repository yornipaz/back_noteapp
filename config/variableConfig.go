package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func loadEnvironments() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file ", err.Error())
		return
	}

}
