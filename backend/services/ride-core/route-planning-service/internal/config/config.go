package config

import "os"

type Config struct {
	Port         string
	RedisAddr    string
	KafkaBrokers string
	OSRMEndpoint string
}

func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		RedisAddr:    getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers: getEnv("KAFKA_BROKERS", "localhost:9092"),
		OSRMEndpoint: getEnv("OSRM_ENDPOINT", "http://localhost:5000"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
