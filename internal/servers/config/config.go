package config

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Config struct {
	Server struct {
		Host  string
		Port  string
		Token string
	}
}

func getStringFromEnv(env string) string {
	res := os.Getenv(env)
	if res == "" {
		panic(fmt.Sprintf("Failed to get env: %s", env))
	}
	return res
}

func NewConfig() *Config {
	config := &Config{}
	config.Server.Host = getStringFromEnv("SERVER_HOST")
	config.Server.Port = getStringFromEnv("SERVER_PORT")

	config.Server.Token = uuid.New().String()

	return config
}
