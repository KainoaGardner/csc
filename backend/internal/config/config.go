package config

import (
	"os"
)

type Config struct {
	PublicHost string
	Port       string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: os.Getenv("PUBLIC_HOST"),
		Port:       os.Getenv("PORT"),
	}
}
