package main

import (
	"encoding/json"
	"io"
	"os"
)

type DBConfig struct {
	Driver string `json:"driver"`
	DSN    string `json:"dsn"`
}

type HermesConfig struct {
	KeyFile  string   `json:"keyfile"`
	DataBase DBConfig `json:"database"`
}

func load_config(configFileName string) (HermesConfig, error) {
	var config HermesConfig

	configFile, err := os.Open(configFileName)
	if err == nil {
		configByte, _ := io.ReadAll(configFile)
		json.Unmarshal(configByte, &config)
	}

	return config, err
}
