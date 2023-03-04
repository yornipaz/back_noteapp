package config

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func loadVariables() {
	err := godotenv.Load("nix-env.env")
	fmt.Println("load variables")
	out, errExec := exec.Command("ls").Output()
	if errExec != nil {
		log.Fatal(errExec)
	}
	fmt.Println(string(out))
	if err != nil {
		log.Fatal("Error loading .env file ", err.Error())
	}

}
