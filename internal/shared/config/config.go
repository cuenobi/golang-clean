package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppName          string
	AppEnv           string
	LogLevel         string
	HTTPAddress      string
	HTTPReadTimeout  int
	HTTPWriteTimeout int

	AuthEnabled      bool
	APIKey           string
	RateLimitEnabled bool
	RateLimitMax     int
	RateLimitWindow  int

	ReadinessDBTimeoutMS int

	PostgresHost    string
	PostgresPort    string
	PostgresUser    string
	PostgresPass    string
	PostgresDB      string
	PostgresSSLMode string

	KafkaBrokers []string

	OutboxPollIntervalMS      int
	OutboxBatchSize           int
	OutboxMaxRetries          int
	OutboxRetryBackoffMS      int
	OutboxProcessingTimeoutMS int
	KafkaPublishTimeoutMS     int

	CircuitBreakerFailures int
	CircuitBreakerOpenMS   int
}

func Load() Config {
	kafkaBrokersRaw := envOrDefault("KAFKA_BROKER", "localhost:9092")
	kafkaBrokers := splitCSV(kafkaBrokersRaw)
	if len(kafkaBrokers) == 0 {
		kafkaBrokers = []string{"localhost:9092"}
	}

	return Config{
		AppName:          envOrDefault("APP_NAME", "golang-clean"),
		AppEnv:           envOrDefault("APP_ENV", "dev"),
		LogLevel:         envOrDefault("LOG_LEVEL", "info"),
		HTTPAddress:      envOrDefault("HTTP_ADDRESS", ":8080"),
		HTTPReadTimeout:  envOrDefaultInt("HTTP_READ_TIMEOUT_SEC", 10),
		HTTPWriteTimeout: envOrDefaultInt("HTTP_WRITE_TIMEOUT_SEC", 15),

		AuthEnabled:      envOrDefaultBool("AUTH_ENABLED", false),
		APIKey:           os.Getenv("API_KEY"),
		RateLimitEnabled: envOrDefaultBool("RATE_LIMIT_ENABLED", true),
		RateLimitMax:     envOrDefaultInt("RATE_LIMIT_MAX", 120),
		RateLimitWindow:  envOrDefaultInt("RATE_LIMIT_WINDOW_SEC", 60),

		ReadinessDBTimeoutMS: envOrDefaultInt("READINESS_DB_TIMEOUT_MS", 1500),

		PostgresHost:    envOrDefault("POSTGRES_HOST", "localhost"),
		PostgresPort:    envOrDefault("POSTGRES_PORT", "5432"),
		PostgresUser:    envOrDefault("POSTGRES_USER", "oms_user"),
		PostgresPass:    envOrDefault("POSTGRES_PASS", "oms_password"),
		PostgresDB:      envOrDefault("POSTGRES_DB", "oms_db"),
		PostgresSSLMode: envOrDefault("POSTGRES_SSL_MODE", "disable"),

		KafkaBrokers: kafkaBrokers,

		OutboxPollIntervalMS:      envOrDefaultInt("OUTBOX_POLL_INTERVAL_MS", 1000),
		OutboxBatchSize:           envOrDefaultInt("OUTBOX_BATCH_SIZE", 50),
		OutboxMaxRetries:          envOrDefaultInt("OUTBOX_MAX_RETRIES", 5),
		OutboxRetryBackoffMS:      envOrDefaultInt("OUTBOX_RETRY_BACKOFF_MS", 500),
		OutboxProcessingTimeoutMS: envOrDefaultInt("OUTBOX_PROCESSING_TIMEOUT_MS", 15000),
		KafkaPublishTimeoutMS:     envOrDefaultInt("KAFKA_PUBLISH_TIMEOUT_MS", 3000),

		CircuitBreakerFailures: envOrDefaultInt("CIRCUIT_BREAKER_FAILURES", 5),
		CircuitBreakerOpenMS:   envOrDefaultInt("CIRCUIT_BREAKER_OPEN_MS", 30000),
	}
}

func (c Config) PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPass,
		c.PostgresDB,
		c.PostgresSSLMode,
	)
}

func (c Config) PostgresMigrationURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.PostgresUser,
		c.PostgresPass,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDB,
		c.PostgresSSLMode,
	)
}

func envOrDefault(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func envOrDefaultInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return parsed
}

func envOrDefaultBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(v)
	if err != nil {
		return fallback
	}
	return parsed
}

func splitCSV(raw string) []string {
	parts := strings.Split(raw, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		clean := strings.TrimSpace(part)
		if clean != "" {
			result = append(result, clean)
		}
	}
	return result
}
