package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jasonlvhit/gocron"
)

func main() {
	loadEnv()
	loadConfig()

	go executeConfigReload()

	log.Print("Starting the service")
	router := Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+envValue("HTTP_PORT", "8000"), router))
}

func executeConfigReload() {
	interval, _ := strconv.Atoi(envValue("CONFIG_RELOAD_INTERVAL", "60"))

	gocron.Every(uint64(interval)).Second().Do(loadConfig)
	<-gocron.Start()
}
