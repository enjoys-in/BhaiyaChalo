package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port            int
	RedisAddr       string
	KafkaBrokers    []string
	OSRMEndpoint    string
	DefaultSpeedKMH float64
}

func Load() *Config {
	return &Config{
		Port:            getEnvInt("PORT", 8080),
		RedisAddr:       getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers:    getEnvSlice("KAFKA_BROKERS", []string{"localhost:9092"}),
		OSRMEndpoint:    getEnv("OSRM_ENDPOINT", "http://localhost:5000"),
		DefaultSpeedKMH: getEnvFloat("DEFAULT_SPEED_KMH", 30.0),
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

func getEnvFloat(key string, fallback float64) float64 {
	if v := os.Getenv(key); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
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
