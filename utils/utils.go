package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	// loads values from .env into the system
	env := os.Getenv("ENV")
	fmt.Println(env)

	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	fmt.Println(".env." + env + ".local")

	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	fmt.Println(".env." + env)

	err := godotenv.Load()
	HandleErr(err)
}
