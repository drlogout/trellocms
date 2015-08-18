package trellocms

import (
	"io/ioutil"
	"os"
	"encoding/json"
)

var (
	configDir = os.Getenv("HOME")
)

type Config struct {
	API     string
	BoardId string
}

func ParseConfig() (Config, error) {

	var config Config
	configFile, err := ioutil.ReadFile(configDir + "/trello-cms-config.json")
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configFile, &config)

	return config, err

}
