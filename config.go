package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const DEFAULT_URL string = "default"
const DEFAULT_URL_FILE string = "default_url.txt"

type RedirectConfig struct {
	name              string
	userAgentPatterns map[string][]string
	urls              map[string][]string
}

var redirectConfig map[string]RedirectConfig

func loadConfig() {

	var configDir = envValue("CONFIG_DIR", "config")
	log.Print(fmt.Sprintf("Loading config from: %v \n", configDir))

	files, err := ioutil.ReadDir(configDir)
	if err != nil {
		log.Fatal(err)
	}
	var newRedirectConfig = map[string]RedirectConfig{}

	for _, f := range files {
		if f.IsDir() {
			fmt.Println("Processing: " + f.Name())
			var config = RedirectConfig{}

			var fileContent = map[string][]string{}

			configFiles, err := ioutil.ReadDir(configDir + "/" + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			for _, config := range configFiles {
				if !config.IsDir() {
					fmt.Println("Processing: " + config.Name())

					file, err := os.Open(configDir + "/" + f.Name() + "/" + config.Name())
					if err != nil {
						log.Fatalf("failed opening file: %s", err)
					}

					scanner := bufio.NewScanner(file)
					scanner.Split(bufio.ScanLines)
					var lines []string

					for scanner.Scan() {
						lines = append(lines, scanner.Text())
					}

					fileContent[config.Name()] = lines
					file.Close()
				}
			}

			log.Print(fmt.Sprintf("fileContent : %v\n", fileContent))

			newRedirectConfig[f.Name()] = config
		}
	}

	log.Print(fmt.Sprintf("Loaded : %v\n", newRedirectConfig))

	redirectConfig = newRedirectConfig
}
