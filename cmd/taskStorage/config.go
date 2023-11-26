package main

import (
	"encoding/json"
	"os"
	"taskStorage/pkg/infrastructure"
)

type Config struct {
	Database infrastructure.DatabaseConfig `json:"database"`
}

func NewConfig() *Config {
	var config Config
	file, err := os.Open("cmd/taskStorage/config/config.json")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
