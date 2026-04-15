package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port            string
	RedisAddr       string
	KafkaBrokers    string
	BatchSize       int
	FlushIntervalMS int
}

func Load() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		RedisAddr:       getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers:    getEnv("KAFKA_BROKERS", "localhost:9092"),
		BatchSize:       getEnvInt("BATCH_SIZE", 100),
		FlushIntervalMS: getEnvInt("FLUSH_INTERVAL_MS", 500),
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
