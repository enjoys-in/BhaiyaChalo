package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port                    int
	RedisAddr               string
	KafkaBrokers            []string
	LocationTTLSeconds      int
	SnapshotIntervalSeconds int
}

func Load() *Config {
	return &Config{
		Port:                    getEnvInt("PORT", 8080),
		RedisAddr:               getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers:            getEnvSlice("KAFKA_BROKERS", []string{"localhost:9092"}),
		LocationTTLSeconds:      getEnvInt("LOCATION_TTL_SECONDS", 300),
		SnapshotIntervalSeconds: getEnvInt("SNAPSHOT_INTERVAL_SECONDS", 5),
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
