package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

// Configuration object to be parsed from files.
type Configuration struct {
	DBConfig DBConfig
}

// DBConfig object to be parsed from files.
type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

//Config returns config after parsing values from env json file.
func Config(env string) Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf("/"+env+".json", &configuration)

	if err != nil {
		fmt.Println("Exception reading config for env " + env)
	}

	return configuration
}
