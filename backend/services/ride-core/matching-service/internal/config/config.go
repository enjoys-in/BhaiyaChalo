package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port             string
	DatabaseURL      string
	RedisAddr        string
	KafkaBrokers     string
	MaxSearchRadiusKM float64
	DefaultRadiusKM  float64
}

func Load() *Config {
	return &Config{
		Port:             getEnv("PORT", "8082"),
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://localhost:5432/matching_db?sslmode=disable"),
		RedisAddr:        getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers:     getEnv("KAFKA_BROKERS", "localhost:9092"),
		MaxSearchRadiusKM: getEnvFloat("MAX_SEARCH_RADIUS_KM", 15.0),
		DefaultRadiusKM:  getEnvFloat("DEFAULT_RADIUS_KM", 5.0),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
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
