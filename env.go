package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

}

func envValue(key string, defaultValue string) string {
	var configDir, exists = os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return configDir
}
