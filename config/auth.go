package config

import (
	"os"
	"time"
)

type AuthConfig struct {
	Secret   string
	ExpHours time.Time
}

func LoadAuthConfig() AuthConfig {

	return AuthConfig{
		Secret:   os.Getenv("AUTH_SECRET"),
		ExpHours: time.Now().Add(48 * time.Hour),
	}
}
