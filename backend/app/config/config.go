package config

import (
	"fmt"
	"os"
)

type Config struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSLMODE  string
	STAGE       string
}

func LoadConfig() (*Config, error) {

	var cfg Config
	var ok bool

	if cfg.DB_HOST, ok = os.LookupEnv("DB_HOST"); !ok {
		return nil, fmt.Errorf("missing DB_HOST")
	}
	if cfg.DB_USER, ok = os.LookupEnv("DB_USER"); !ok {
		return nil, fmt.Errorf("missing DB_USER")
	}
	if cfg.DB_PASSWORD, ok = os.LookupEnv("DB_PASSWORD"); !ok {
		return nil, fmt.Errorf("missing DB_PASSWORD")
	}
	if cfg.DB_NAME, ok = os.LookupEnv("DB_NAME"); !ok {
		return nil, fmt.Errorf("missing DB_NAME")
	}
	if cfg.DB_PORT, ok = os.LookupEnv("DB_PORT"); !ok {
		return nil, fmt.Errorf("missing DB_PORT")
	}
	if cfg.DB_SSLMODE, ok = os.LookupEnv("DB_SSLMODE"); !ok {
		return nil, fmt.Errorf("missing DB_SSLMODE")
	}
	if cfg.STAGE, ok = os.LookupEnv("STAGE"); !ok {
		return nil, fmt.Errorf("missing STAGE")
	}

	return &cfg, nil
}
