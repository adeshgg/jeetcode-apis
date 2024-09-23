package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBUser        string `json:"DBUser"`
	DBPassword    string `json:"DBPassword"`
	DBName        string `json:"DBName"`
	DBHost        string `json:"DBHost"`
	DBPort        string `json:"DBPort"`
	RedisAddress  string `json:"RedisAddress"`
	RedisPassword string `json:"RedisPassword"`
}

func LoadConfig(configFilePath string) (*Config, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
