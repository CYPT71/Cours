package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var Config *Configuration

type Configuration struct {
	Server ServerConfiguration `yaml:"server"`
	Mysql  MysqlConfiguration  `yaml:"mysql"`
}

type ServerConfiguration struct {
	Port      string `yaml:"port"`
	Address   string `yaml:"address"`
	SecretKey string `yaml:"secretKey"`
}

type MysqlConfiguration struct {
	Dns string `yaml:"dns"`
}

// Server,Repository initialize configuration
func Setup() {
	var config *Configuration
	configFile, err := os.Open("../BackEnd/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	if err = decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding config file, %s", err)
	}
	Config = config
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
