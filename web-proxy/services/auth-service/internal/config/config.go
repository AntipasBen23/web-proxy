package config

import (
	"log"
	"os"
)

type Config struct {
	Port           string
	AuthHeaderName string
	AuthSecretKey  string
	AuthStrategy   string
}

func Load() *Config {
	cfg := &Config{
		Port:           getEnv("PORT", "8081"),
		AuthHeaderName: getEnv("AUTH_HEADER_NAME", "Authorization"),
		AuthSecretKey:  os.Getenv("AUTH_SECRET_KEY"), 
		AuthStrategy:   getEnv("AUTH_STRATEGY", "api_key"),
	}

	if cfg.AuthSecretKey == "" {
		log.Fatal("❌ AUTH_SECRET_KEY is required but not set")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
