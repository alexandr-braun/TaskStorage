package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database `json:"database"`
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
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
