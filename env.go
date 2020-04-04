package main

import (
	"log"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// value, exists := os.LookupEnv("qwe")
	// log.Print(fmt.Sprintf("Env : %v %v\n", value, exists))

}
