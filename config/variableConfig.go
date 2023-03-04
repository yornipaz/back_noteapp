package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func loadVariables() {
	err := godotenv.Load("nix-env.env")
	fmt.Println("load variables")
	if err != nil {
		log.Fatal("Error loading .env file ", err.Error())
	}

}
