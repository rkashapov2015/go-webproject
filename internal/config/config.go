package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(filePath string) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
