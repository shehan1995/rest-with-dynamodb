package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server ServerConfig
	AWS    AWSConfig
}

type ServerConfig struct {
	Port         int `json:"port"`
	IdleTimeout  int `json:"idle_timeout"`
	ReadTimeout  int `json:"read_timeout"`
	WriteTimeout int `json:"write_timeout"`
}

type AWSConfig struct {
	AccessKey     string `json:"access_key"`
	SecretKey     string `json:"secret_key"`
	Endpoint      string `json:"endpoint"`
	Region        string `json:"region"`
	DynamoDBTable string `json:"dynamodb_table"`
}

func LoadConfig() (Config, error) {

	f, err := os.ReadFile("config/config.json")
	if err != nil {
		return Config{}, err
	}

	conf := Config{}

	err = json.Unmarshal(f, &conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}
