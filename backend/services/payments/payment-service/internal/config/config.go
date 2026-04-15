package config

import (
	"os"
	"strings"
)

type Config struct {
	Port              string
	DatabaseURL       string
	RedisAddr         string
	KafkaBrokers      []string
	PaymentGatewayURL string
	PaymentGatewayKey string
}

func Load() *Config {
	return &Config{
		Port:              getEnv("PORT", "8080"),
		DatabaseURL:       getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/payment_service?sslmode=disable"),
		RedisAddr:         getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers:      strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ","),
		PaymentGatewayURL: getEnv("PAYMENT_GATEWAY_URL", "https://api.razorpay.com"),
		PaymentGatewayKey: getEnv("PAYMENT_GATEWAY_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
