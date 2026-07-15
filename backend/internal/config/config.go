package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Env                 string
	Port                string
	DatabaseURL         string
	JWTAccessSecret     string
	JWTRefreshSecret    string
	AccessTokenTTLMin   int
	RefreshTokenTTLDays int
	AllowedOrigins      []string
}

func Load() (Config, error) {
	cfg := Config{
		Env:                 getString("APP_ENV", "development"),
		Port:                getString("PORT", "8080"),
		DatabaseURL:         getString("DATABASE_URL", ""),
		JWTAccessSecret:     getString("JWT_ACCESS_SECRET", "dev-access-secret-change-me"),
		JWTRefreshSecret:    getString("JWT_REFRESH_SECRET", "dev-refresh-secret-change-me"),
		AccessTokenTTLMin:   getInt("ACCESS_TOKEN_TTL_MINUTES", 15),
		RefreshTokenTTLDays: getInt("REFRESH_TOKEN_TTL_DAYS", 30),
		AllowedOrigins:      getSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000"}),
	}

	if cfg.Port == "" {
		return Config{}, fmt.Errorf("PORT is required")
	}

	return cfg, nil
}

func getString(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}

	return fallback
}

func getInt(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}

func getSlice(key string, fallback []string) []string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	parts := strings.Split(value, ",")
	items := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			items = append(items, trimmed)
		}
	}

	if len(items) == 0 {
		return fallback
	}

	return items
}