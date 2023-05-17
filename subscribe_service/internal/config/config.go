package config

import (
	"encoding/json"
	"net"
	"os"
)

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func (h *HTTP) GetAddress() string {
	return net.JoinHostPort(h.Host, h.Port)
}

type DB struct {
	DSN  string `json:"dsn"`
	USER string `json:"user"`
	DB   string `json:"db"`
	PASS string `json:"password"`
}

type Config struct {
	HTTP HTTP `json:"http"`
	DB   DB   `json:"db"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetDBConfig() *DB {

	return &c.DB
}
