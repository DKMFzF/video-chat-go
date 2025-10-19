package config

import (
	"os"

	gotenv "github.com/subosito/gotenv"
)

type Config struct {
	Port string
}

func Load() *Config {
	_ = gotenv.Load()

	return &Config{
		Port: os.Getenv("PORT"),
	}
}
