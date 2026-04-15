package config

import "os"

type Config struct {
	Port           int
	DatabaseURL    string
	RedisAddr      string
	KafkaBrokers   []string
	JWTSecret      string
	OTPServiceAddr string
}

func Load() *Config {
	return &Config{
		Port:           getEnvInt("PORT", 8080),
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://localhost:5432/auth"),
		RedisAddr:      getEnv("REDIS_ADDR", "localhost:6379"),
		KafkaBrokers:   []string{getEnv("KAFKA_BROKERS", "localhost:9092")},
		JWTSecret:      getEnv("JWT_SECRET", "change-me"),
		OTPServiceAddr: getEnv("OTP_SERVICE_ADDR", "localhost:9001"),
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
	i := 0
	for _, c := range v {
		i = i*10 + int(c-'0')
	}
	return i
}
