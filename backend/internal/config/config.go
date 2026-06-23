package config

import "os"

type Config struct {
	Addr         string
	DatabasePath string
}

func Load() *Config {
	return &Config{
		Addr:         getEnv("SERVER_ADDR", ":8080"),
		DatabasePath: getEnv("DATABASE_PATH", "./data/media-hub.db"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
