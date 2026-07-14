package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func (d DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode,
	)
}

type ServerConfig struct {
	Port string
	Env  string
}

type JWTConfig struct {
	Secret     string
	Expiration time.Duration
}

type Config struct {
	Database   DatabaseConfig
	Server     ServerConfig
	JWT        JWTConfig
	CORSOrigin string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	expiration, err := time.ParseDuration(getEnv("JWT_EXPIRATION", "24h"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRATION: %w", err)
	}

	cfg := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "smartkiki_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Env:  getEnv("SERVER_ENV", "development"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", ""),
			Expiration: expiration,
		},
		CORSOrigin: getEnv("CORS_ORIGIN", "http://localhost:5173"),
	}

	if cfg.JWT.Secret == "" {
		return nil, fmt.Errorf("JWT_SECRET must be set")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
