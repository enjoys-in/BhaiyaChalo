package config

import "os"

type Config struct {
	Port               string
	DatabaseURL        string
	RedisAddr          string
	KafkaBrokers       string
	ETAServiceAddr     string
	PricingServiceAddr string
}

func Load() *Config {
	return &Config{
		Port:               getEnv("PORT", "8081"),
		DatabaseURL:        getEnv("DATABASE_URL", "postgres://localhost:5432/search_db?sslmode=disable"),
		RedisAddr:          getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers:       getEnv("KAFKA_BROKERS", "localhost:9092"),
		ETAServiceAddr:     getEnv("ETA_SERVICE_ADDR", "localhost:50053"),
		PricingServiceAddr: getEnv("PRICING_SERVICE_ADDR", "localhost:50051"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
