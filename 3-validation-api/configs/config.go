// Package configs structs
package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB    DBConfig
	Auth  AuthConfig
	Email EmailConfig
}

type DBConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type EmailConfig struct {
	Email   string
	Pass    string
	Address string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			log.Println("ENV_FILE_ERROR, using default config")
		}
	}
	return &Config{
		DB: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
		Email: EmailConfig{
			Email:   os.Getenv("EMAIL"),
			Pass:    os.Getenv("PASSWORD"),
			Address: os.Getenv("ADDRESS"),
		},
	}
}
