package config

import "os"

type Config struct {
	Port           int
	RedisAddr      string
	AuthServiceAddr string
}

func Load() *Config {
	return &Config{
		Port:            getEnvInt("PORT", 3001),
		RedisAddr:       getEnv("REDIS_ADDR", "localhost:6379"),
		AuthServiceAddr: getEnv("AUTH_SERVICE_ADDR", "localhost:8081"),
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
