package config

import (
	"os"
)

type AuthConfig struct {
	Secret string
}

func LoadAuthConfig() AuthConfig {

	return AuthConfig{
		Secret: os.Getenv("AUTH_SECRET"),
	}
}
