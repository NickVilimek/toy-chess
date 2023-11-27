package scripts

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type DbConfig struct {
	MongoConnectionString string `json:"mongo_connection_string"`
}

var DbConfigInstance *DbConfig

func InitToolingConfig() {

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Can't get working directory")
	}

	filePath := filepath.Join(currDir, "db-tooling", "db-config.json")

	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Couldn't Load Config File")
	}

	defer configFile.Close()

	var tempConfig *DbConfig

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&tempConfig); err != nil {
		log.Fatal("Error parsing the json from the config file")
	}

	DbConfigInstance = tempConfig
}
