package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppName      string
	AppEnv       string
	LogLevel     string
	HTTPAddress  string
	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPass string
	PostgresDB   string
	KafkaBrokers []string
}

func Load() Config {
	return Config{
		AppName:      envOrDefault("APP_NAME", "golang-clean"),
		AppEnv:       envOrDefault("APP_ENV", "dev"),
		LogLevel:     envOrDefault("LOG_LEVEL", "info"),
		HTTPAddress:  envOrDefault("HTTP_ADDRESS", ":8080"),
		PostgresHost: envOrDefault("POSTGRES_HOST", "localhost"),
		PostgresPort: envOrDefault("POSTGRES_PORT", "5432"),
		PostgresUser: envOrDefault("POSTGRES_USER", "oms_user"),
		PostgresPass: envOrDefault("POSTGRES_PASS", "oms_password"),
		PostgresDB:   envOrDefault("POSTGRES_DB", "oms_db"),
		KafkaBrokers: []string{envOrDefault("KAFKA_BROKER", "localhost:9092")},
	}
}

func (c Config) PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPass,
		c.PostgresDB,
	)
}

func (c Config) PostgresMigrationURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.PostgresUser,
		c.PostgresPass,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDB,
	)
}

func envOrDefault(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
