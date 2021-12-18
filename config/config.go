package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() (config Config) {
	// We intentionally ignore the error here because there will not be a .env file when
	// the application is deployed
	godotenv.Load()

	// Expand all environment variables
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		os.Setenv(kv[0], os.ExpandEnv(kv[1]))
	}

	return Config{
		Port: os.Getenv("PORT"),
	}
}
