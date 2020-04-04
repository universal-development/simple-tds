package main

import (
	"fmt"
	"log"
	"os"
)

func loadConfig() {

	var configDir = envValue("CONFIG_DIR", "config")

	log.Print(fmt.Sprintf("Loading config from: %v \n", configDir))

}

func envValue(key string, defaultValue string) string {
	var configDir, exists = os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return configDir
}
