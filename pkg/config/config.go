package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Mode         string `env:"MODE""`
	DbUser       string `env:"POSTGRES_USER"`
	DbPass       string `env:"POSTGRES_PASSWORD"`
	DbHost       string `env:"DB_HOST"`
	DbName       string `env:"POSTGRES_DB"`
	ClientId     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
}

func NewConfig() (*Configuration, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("No .env file could be found %v", err)
	}

	cfg := Configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
