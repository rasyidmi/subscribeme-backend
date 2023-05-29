package config

import (
	"os"
)

type SiakNGConfig struct {
	Username string
	Password string
}

func LoadSiakNGConfig() SiakNGConfig {

	return SiakNGConfig{
		Username: os.Getenv("SIAKNG_USERNAME"),
		Password: os.Getenv("SIAKNG_PASSWORD"),
	}
}
