package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	EngineAddr string
	LogicAddr  string

	// Auth / Cookie
	SessionSecret  string // used to namespace/rotate sessions (not strictly required for opaque tokens but good to have)
	CookieName     string
	CookieDomain   string
	CookieSecure   bool
	CookieMaxAge   int // seconds
	DBDSN          string
	RedisAddr      string
	SessionBackend string // "redis" | "memory"
}

func get(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getBool(key string, def bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}
	return b
}

func getInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return i
}

func Load() Config {
	// Try to guess cookie domain if provided like "api.localhost"
	domain := strings.TrimSpace(os.Getenv("COOKIE_DOMAIN"))

	return Config{
		// Defaults that work nicely inside Docker Compose; override on host
		EngineAddr: get("ENGINE_ADDR", "localhost:9101"),
		LogicAddr:  get("LOGIC_ADDR", "localhost:9002"),

		SessionSecret:  get("SESSION_SECRET", "dev-secret-change-me"),
		CookieName:     get("COOKIE_NAME", "harmonia_session"),
		CookieDomain:   domain,
		CookieSecure:   getBool("COOKIE_SECURE", false), // set true in prod/https
		CookieMaxAge:   getInt("COOKIE_MAX_AGE", 7*24*3600),
		DBDSN:          get("DB_DSN", "harmonia:harmonia@tcp(localhost:3306)/harmonia?parseTime=true"),
		RedisAddr:      get("REDIS_ADDR", "localhost:6379"),
		SessionBackend: get("SESSION_BACKEND", "redis"),
	}
}
