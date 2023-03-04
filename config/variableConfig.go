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
	out, errCd := exec.Command("cd out/").Output()
	if errCd != nil {
		log.Fatal(errCd)
	}
	fmt.Println(string(out))
	ls, errLs := exec.Command("cd out/").Output()
	if errLs != nil {
		log.Fatal(errLs)
	}
	fmt.Println(string(ls))
	if err != nil {
		log.Fatal("Error loading .env file ", err.Error())
	}

}
