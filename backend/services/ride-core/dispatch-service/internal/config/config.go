package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port                int
	DatabaseURL         string
	RedisAddr           string
	KafkaBrokers        []string
	OfferTimeoutSeconds int
	MaxRetries          int
}

func Load() *Config {
	return &Config{
		Port:                getEnvInt("PORT", 8080),
		DatabaseURL:         getEnv("DATABASE_URL", "postgres://localhost:5432/dispatch"),
		RedisAddr:           getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers:        getEnvSlice("KAFKA_BROKERS", []string{"localhost:9092"}),
		OfferTimeoutSeconds: getEnvInt("OFFER_TIMEOUT_SECONDS", 30),
		MaxRetries:          getEnvInt("MAX_RETRIES", 3),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func getEnvSlice(key string, fallback []string) []string {
	if v := os.Getenv(key); v != "" {
		return strings.Split(v, ",")
	}
	return fallback
}
