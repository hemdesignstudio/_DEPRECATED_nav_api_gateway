package config

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Configuration struct {
	BaseUrl         string
	CompanyEndpoint string
	Passwd          string
	Username        string
	Host            string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	path := "./config/config.json"
	err := gonfig.GetConf(path, &configuration)
	if err != nil {
		log.Printf("unable to find config file in the following path: %s", path)
	}
	return configuration
}

