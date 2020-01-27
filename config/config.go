package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// config path
const (
	c = "config.json"
)

// Config struct
type Config struct {
	Host string
	Port string
}

// NewConfig constructor
func NewConfig() (*Config, error) {
	data, err := ioutil.ReadFile(c)

	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = json.Unmarshal(data, cfg)

	if err != nil {
		return nil, err
	}

	if h := os.Getenv("HOST"); h != "" {
		cfg.Host = h
	}

	if p := os.Getenv("PORT"); p != "" {
		cfg.Port = p
	}

	return cfg, nil
}
