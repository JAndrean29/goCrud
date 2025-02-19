package config

import env "github.com/caarlos0/env/v11"

type Cfg struct {
	Port        string
	Database    DbConfig
	ServiceName string
}

type DbConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SSLMODE  string
}

func LoadConfig() (*Cfg, error) {
	cfg := &Cfg{}

	if err := env.Parse(cfg); err != nil {
		return &Cfg{}, err
	}

	return cfg, nil
}
