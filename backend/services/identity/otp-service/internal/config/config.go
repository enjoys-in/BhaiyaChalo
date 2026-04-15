package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port            string
	DatabaseURL     string
	RedisAddr       string
	KafkaBrokers    []string
	OTPLength       int
	OTPExpirySeconds int
	SMSProviderURL  string
}

func Load() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		DatabaseURL:     getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/otp_service?sslmode=disable"),
		RedisAddr:       getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers:    strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ","),
		OTPLength:       getEnvInt("OTP_LENGTH", 6),
		OTPExpirySeconds: getEnvInt("OTP_EXPIRY_SECONDS", 300),
		SMSProviderURL:  getEnv("SMS_PROVIDER_URL", "http://localhost:9090/sms"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
