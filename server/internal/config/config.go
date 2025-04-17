package config

import (
	"os"
)

type Config struct {
	Port      string
	JWTSecret string
	DBURL     string
	// Add other configuration fields as needed
}

func Load() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
		DBURL:     getEnv("DB_URL", "mongodb://localhost:27017"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
