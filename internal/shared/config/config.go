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
	AppTimeZone      string
	LogLevel         string
	HTTPAddress      string
	HTTPReadTimeout  int
	HTTPWriteTimeout int
	HTTPBodyLimitMB  int

	AuthEnabled      bool
	APIKey           string
	RateLimitEnabled bool
	RateLimitMax     int
	RateLimitWindow  int
	CORSOrigins      []string
	CORSMethods      []string
	CORSHeaders      []string
	CORSExpose       []string
	CORSAllowCreds   bool
	CORSMaxAgeSec    int

	ReadinessDBTimeoutMS int

	PostgresHost    string
	PostgresPort    string
	PostgresUser    string
	PostgresPass    string
	PostgresDB      string
	PostgresSSLMode string
	PostgresTZ      string

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
		AppTimeZone:      envOrDefault("APP_TIMEZONE", "UTC"),
		LogLevel:         envOrDefault("LOG_LEVEL", "info"),
		HTTPAddress:      envOrDefault("HTTP_ADDRESS", ":8080"),
		HTTPReadTimeout:  envOrDefaultInt("HTTP_READ_TIMEOUT_SEC", 10),
		HTTPWriteTimeout: envOrDefaultInt("HTTP_WRITE_TIMEOUT_SEC", 15),
		HTTPBodyLimitMB:  envOrDefaultInt("HTTP_BODY_LIMIT_MB", 4),

		AuthEnabled:      envOrDefaultBool("AUTH_ENABLED", false),
		APIKey:           os.Getenv("API_KEY"),
		RateLimitEnabled: envOrDefaultBool("RATE_LIMIT_ENABLED", true),
		RateLimitMax:     envOrDefaultInt("RATE_LIMIT_MAX", 120),
		RateLimitWindow:  envOrDefaultInt("RATE_LIMIT_WINDOW_SEC", 60),
		CORSOrigins: splitCSV(envOrDefault(
			"CORS_ALLOWED_ORIGINS",
			"http://localhost:3000,http://localhost:5173",
		)),
		CORSMethods: splitCSV(envOrDefault(
			"CORS_ALLOWED_METHODS",
			"GET,POST,PUT,PATCH,DELETE,OPTIONS",
		)),
		CORSHeaders: splitCSV(envOrDefault(
			"CORS_ALLOWED_HEADERS",
			"Origin,Content-Type,Accept,Authorization,X-API-Key,X-Request-ID,Idempotency-Key",
		)),
		CORSExpose: splitCSV(envOrDefault(
			"CORS_EXPOSE_HEADERS",
			"X-Request-ID",
		)),
		CORSAllowCreds: envOrDefaultBool("CORS_ALLOW_CREDENTIALS", false),
		CORSMaxAgeSec:  envOrDefaultInt("CORS_MAX_AGE_SEC", 300),

		ReadinessDBTimeoutMS: envOrDefaultInt("READINESS_DB_TIMEOUT_MS", 1500),

		PostgresHost:    envOrDefault("POSTGRES_HOST", "localhost"),
		PostgresPort:    envOrDefault("POSTGRES_PORT", "5432"),
		PostgresUser:    envOrDefault("POSTGRES_USER", "oms_user"),
		PostgresPass:    envOrDefault("POSTGRES_PASS", "oms_password"),
		PostgresDB:      envOrDefault("POSTGRES_DB", "oms_db"),
		PostgresSSLMode: envOrDefault("POSTGRES_SSL_MODE", "disable"),
		PostgresTZ:      envOrDefault("POSTGRES_TIMEZONE", "UTC"),

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
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPass,
		c.PostgresDB,
		c.PostgresSSLMode,
		c.PostgresTZ,
	)
}

func (c Config) PostgresMigrationURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&timezone=%s",
		c.PostgresUser,
		c.PostgresPass,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDB,
		c.PostgresSSLMode,
		c.PostgresTZ,
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
