package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBSchema   string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBSchema:   os.Getenv("DB_SCHEMA"),
	}

	if cfg.Port == "" {
		return nil, fmt.Errorf("missing PORT environment variable")
	}

	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBSchema == "" {
		return nil, fmt.Errorf("missing database configuration")
	}

	return cfg, nil
}
