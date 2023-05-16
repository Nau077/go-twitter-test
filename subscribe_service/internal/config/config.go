package config

import (
	"encoding/json"
	"net"
	"os"
)

const (
	dbPassEscSeq = "{password}"
	password     = "note-service-password"
)

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func (h *HTTP) GetAddress() string {
	return net.JoinHostPort(h.Host, h.Port)
}

type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

type Config struct {
	HTTP HTTP `json:"http"`
	DB   DB   `json:"db"`
}

// new config
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

// // Get db config
// func (c *Config) GetDBConfig() (*pgxpool.Config, error) {
// 	dbDsn := strings.ReplaceAll(c.DB.DSN, dbPassEscSeq, password)

// 	poolConfig, err := pgxpool.ParseConfig(dbDsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	poolConfig.ConnConfig.BuildStatementCache = nil
// 	poolConfig.ConnConfig.PreferSimpleProtocol = true
// 	poolConfig.MaxConns = c.DB.MaxOpenConnections

// 	return poolConfig, nil
// }
