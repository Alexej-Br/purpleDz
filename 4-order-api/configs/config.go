// Package configs
package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DBConfig
	JWT JWTConfig
}

type DBConfig struct {
	DSN string
}

type JWTConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load("../.env")
		if err != nil {
			fmt.Printf("ENV_FILE_ERROR, using default config: %v\n", err)
		}
	}
	return &Config{
		DB: DBConfig{
			DSN: os.Getenv("DSN"),
		},
		JWT: JWTConfig{
			Secret: os.Getenv("SECRET"),
		},
	}
}
