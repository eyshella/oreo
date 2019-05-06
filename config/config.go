package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Accuracy          float64
	Steps             int
	NumberOfVerticles int
	ResultPath        string
	LogsEnabled       bool
	Alpha             string
	GammaParam        float64
}

var Config *Configuration

func SetUpConfig(path string) {
	log.Printf("OpenConfigFile started. path: %s", path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("OpenConfigFile Error: %s", err.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = &Configuration{}
	err = decoder.Decode(Config)
	if err != nil {
		log.Fatalf("OpenConfigFile Error: %s", err.Error())
	}
	log.Printf("OpenConfigFile finished. Config: %v", Config)
}
