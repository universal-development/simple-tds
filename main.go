package main

import (
	"log"
	"net/http"
)

func main() {
	loadEnv()
	loadConfig()

	log.Print("Starting the service")
	router := Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+envValue("HTTP_PORT", "8000"), router))
}
