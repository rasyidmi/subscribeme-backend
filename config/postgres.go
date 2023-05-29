package config

import (
	"os"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func LoadPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}
