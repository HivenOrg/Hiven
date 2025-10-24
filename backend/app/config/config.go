package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	STAGE                  string // "dev", "test", "prod"
	AWS_REGION             string
	AWS_ACCESS_KEY_ID      string // Don't set in prod, will be infered from IAM
	AWS_SECRET_ACCESS_KEY  string // Don't set in prod, will be infered from IAM
	S3_BUCKET_NAME         string
	DB_HOST                string
	DB_USER                string
	DB_PASSWORD            string
	DB_NAME                string
	DB_PORT                string
	DB_SSLMODE             string
	JWT_SECRET_KEY         string
	TOKEN_EXPIRES_IN_HOURS int
}

func LoadConfig() (*Config, error) {

	var cfg Config
	var ok bool

	if cfg.STAGE, ok = os.LookupEnv("STAGE"); !ok {
		return nil, fmt.Errorf("missing STAGE")
	}

	if cfg.AWS_REGION, ok = os.LookupEnv("AWS_REGION"); !ok {
		return nil, fmt.Errorf("missing AWS_REGION")
	}

	// Prod: leave blank so that AWS SDK can pick these up from IAM role automatically
	if cfg.STAGE == "dev" || cfg.STAGE == "test" {
		if cfg.AWS_ACCESS_KEY_ID, ok = os.LookupEnv("AWS_ACCESS_KEY_ID"); !ok {
			return nil, fmt.Errorf("missing AWS_ACCESS_KEY_ID")
		}
		if cfg.AWS_SECRET_ACCESS_KEY, ok = os.LookupEnv("AWS_SECRET_ACCESS_KEY"); !ok {
			return nil, fmt.Errorf("missing AWS_SECRET_ACCESS_KEY")
		}
	}

	if cfg.S3_BUCKET_NAME, ok = os.LookupEnv("S3_BUCKET_NAME"); !ok {
		return nil, fmt.Errorf("missing S3_BUCKET_NAME")
	}

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

	if cfg.JWT_SECRET_KEY, ok = os.LookupEnv("JWT_SECRET_KEY"); !ok {
		return nil, fmt.Errorf("missing JWT_SECRET_KEY")
	}

	expires_string, ok := os.LookupEnv("TOKEN_EXPIRES_IN_HOURS")
	if !ok {
		return nil, fmt.Errorf("missing TOKEN_EXPIRES_IN_HOURS")
	}
	var err error
	cfg.TOKEN_EXPIRES_IN_HOURS, err = strconv.Atoi(expires_string)
	if err != nil {
		return nil, fmt.Errorf("error converting string to int")
	}

	return &cfg, nil
}
