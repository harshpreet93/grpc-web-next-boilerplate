package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

// Config stores the config needed by the app to operate
type Config struct {
	DBPass string `json:"DBPass"`
	DBHost string `json:"DBHost"`
}

var configLocations = map[string]string{
	"prod":        "/etc/config/prod.json",
	"test":        "./test.json",
	"testInvalid": "./test_invalid.json",
}

var defaultConfig = &Config{DBPass: "devpass", DBHost: "localhost:3306"}

// New creates a Config object for the env specified
func New(env string) (Config, error) {
	location, found := configLocations[env]
	if !found {
		log.Printf("Cannot find config file %s, falling back to default config", location)
		return *defaultConfig, nil
	}
	config, err := createConfigFrom(location)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func createConfigFrom(file string) (Config, error) {
	configJSON, err := ioutil.ReadFile(file)

	if err != nil {
		return Config{}, errors.Errorf("error reading config file %+v", err)
	}
	config := Config{}
	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
