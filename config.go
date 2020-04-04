package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

const AGENT_NAME = "_agent.txt"
const URL_NAME = "_url.txt"
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
			var configName = f.Name()
			var config = RedirectConfig{}

			config.name = configName
			config.urls = map[string][]string{}
			config.userAgentPatterns = map[string][]string{}

			newRedirectConfig[configName] = config

			// loading file content in a map
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
			// extract each
			for key, element := range fileContent {
				if strings.HasSuffix(key, AGENT_NAME) {
					blockName := strings.ReplaceAll(key, AGENT_NAME, "")
					blockUrlsFile := blockName + URL_NAME
					config.name = blockName
					config.urls[blockName] = fileContent[blockUrlsFile]
					config.userAgentPatterns[blockName] = element
				}
			}
			config.urls[DEFAULT_URL_FILE] = fileContent[DEFAULT_URL_FILE]
		}
	}

	log.Print(fmt.Sprintf("Loaded config : %v\n", newRedirectConfig))

	redirectConfig = newRedirectConfig
}

func resolve(name string, agent string) string {

	config, ok := redirectConfig[name]
	if !ok {
		log.Print("No config for name " + name)
		return ""
	}

	var matchedKey = DEFAULT_URL_FILE

	for key, patterns := range config.userAgentPatterns {
		for _, pattern := range patterns {
			match, _ := regexp.MatchString(pattern, agent)
			if match {
				matchedKey = key
				break
			}
		}
		if matchedKey != DEFAULT_URL_FILE {
			break
		}
	}
	log.Print("Matched group for " + name + " / " + agent + " = " + matchedKey)
	var urls, foundPattern = config.urls[matchedKey]

	if !foundPattern || len(urls) == 0 {
		return ""
	}

	rand.Seed(time.Now().Unix())
	var url = urls[rand.Intn(len(urls))]

	return url
}
