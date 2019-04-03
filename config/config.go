package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config - Represents database server and credentials
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
