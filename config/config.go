package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	serverPort  int32
	refreshRate string
	Endpoints   []EndPoint `json:"endpoints"`
}

type EndPoint struct {
	URL                string `json:"url"`
	ContentRequirement string `json:"content_requirement,omitempty"`
}

// LoadConfigFromFile loads the configuration file given a filepath
func LoadConfigFromFile(filePath string) (*Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", filePath, err)
	}

	conf, err := loadConfig(f)

	if err != nil {
		return nil, fmt.Errorf("error retrieving log file: %v", err)
	}

	return conf, nil
}

func loadConfig(file io.ReadWriteCloser) (*Config, error) {
	var c Config
	json.NewDecoder(file).Decode(&c)
	return &c, nil
}

func validateConfig(c Config) error {
	return fmt.Errorf("function not implemented yet")
}
