package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Cfg struct {
	Port        string   `env:"PORT" envDefault:"8080"`
	Database    DbConfig `envPrefix:"DB_"`
	ServiceName string   `env:"SERVICE_NAME"`
}

type DbConfig struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT" envDefault:"5432"`
	Database string `env:"NAME"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	SSLMODE  string `env:"SSL_MODE" envdefault:"disable"`
}

func LoadConfig() (*Cfg, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found, relying on system environment variables.")
	}
	cfg := &Cfg{}

	if err := env.Parse(cfg); err != nil {
		fmt.Println(cfg)
		return &Cfg{}, err
	}

	fmt.Println(cfg)
	return cfg, nil
}
