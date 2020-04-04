package main

import (
	"log"
	"net/http"
)

func main() {
	loadEnv()
	log.Print("Starting the service")
	router := Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
