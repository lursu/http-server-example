package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config top level configuration structure
type Config struct {
	Port int
	DataBase
}

// DataBase contains all of the configuration options needs for a new db connection to be made
type DataBase struct {
	Host     string `envconfig:"DB_HOST"`
	Port     int    `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
	SSLMode  string `encconfig:"DB_SSLMODE" default:"disable"`
}

// Reads the env vars and parses them into the parent config struct
func Read(name string) (Config, error) {
	var cfg Config
	return cfg, envconfig.Process(name, &cfg)
}

// Url helper to format the connection information into a connection Url
func (d DataBase) Url() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}
