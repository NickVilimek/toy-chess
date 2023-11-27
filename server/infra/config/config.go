package infra

import (
	"encoding/json"
	"log"
	"os"
)

var GlobalConfigInstance *Config

func InitGlobalConfig(configLocation string) {
	configFile, err := os.Open(configLocation)
	if err != nil {
		log.Fatal("Couldn't Load Config File")
	}

	defer configFile.Close()

	var config *Config
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Fatal("Error parsing the json from the config file")
	}

	GlobalConfigInstance = config
}
