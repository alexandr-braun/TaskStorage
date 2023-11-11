package main

import (
	"encoding/json"
	"os"
	"taskStorage/pkg/infrastructure"
)

type Config struct {
	Database infrastructure.DatabaseConfig `json:"database"`
}

func NewConfig() (*Config, error) {
	var config Config
	file, err := os.Open("cmd/taskStorage/config/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
