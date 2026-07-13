package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL             string
	LimenSecret       string
	ApiVersion        string
	ApiPort           string
	GooseDriver       string
	GooseDBString     string
	GooseMigrationDir string
	GooseTable        string
	CorsAllowOrigin   string
}

func NewConfig(envPath string) (*Config, error) {
	if err := godotenv.Load(envPath); err != nil {
		return nil, err
	}

	cfg := &Config{
		DBURL:       os.Getenv("DB_URL"),
		LimenSecret: os.Getenv("LIMEN_SECRET"),
		ApiVersion:  os.Getenv("API_VERSION"),
		ApiPort:     os.Getenv("API_PORT"),

		GooseDriver:       os.Getenv("GOOSE_DRIVER"),
		GooseDBString:     os.Getenv("GOOSE_DBSTRING"),
		GooseMigrationDir: os.Getenv("GOOSE_MIGRATION_DIR"),
		GooseTable:        os.Getenv("GOOSE_TABLE"),

		CorsAllowOrigin: os.Getenv("CORS_ALLOWED_ORIGIN"),
	}

	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is required")
	}

	if cfg.LimenSecret == "" {
		return nil, fmt.Errorf("LIMEN_SECRET is required")
	}

	if cfg.ApiVersion == "" {
		return nil, fmt.Errorf("API_VERSION is required")
	}

	if cfg.ApiPort == "" {
		return nil, fmt.Errorf("API_PORT is required")
	}

	if cfg.GooseDriver == "" {
		return nil, fmt.Errorf("GOOSE_DRIVER is required")
	}

	if cfg.GooseDBString == "" {
		return nil, fmt.Errorf("GOOSE_DBSTRING is required")
	}

	if cfg.GooseMigrationDir == "" {
		return nil, fmt.Errorf("GOOSE_MIGRATION_DIR is required")
	}

	if cfg.CorsAllowOrigin == "" {
		cfg.CorsAllowOrigin = "http://localhost:5173"
	}

	return cfg, nil
}
