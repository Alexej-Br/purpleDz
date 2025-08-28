// Package configs
package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	DSN string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("ENV_FILE_ERROR, using default config: %v\n", err)
	}
	return &Config{
		DB: DBConfig{
			DSN: os.Getenv("DSN"),
		},
	}
}
