package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       int
	Host       string
	Debug      bool
	DBPath     string
	TestingRun bool
	EmailerKey string
	JWT_Secret string
	App_Name   string
	LogDir     string
}

var AppConfig *Config

// init autoruns
func init() {
	godotenv.Load()

	AppConfig = &Config{
		Port:       getEnvInt("PORT", 8080),
		Host:       getEnvString("HOST", "localhost"),
		Debug:      getEnvBool("DEBUG", false), //currently not in use
		DBPath:     getEnvString("DB_DIR", "/var/lib/AuthBase/db"),
		LogDir:     getEnvString("LOG_DIR", "/var/log/AuthBase"),
		TestingRun: getEnvBool("TESTING", false),
		EmailerKey: getEnvString("EMAILER_SECRET_KEY", ""),
		App_Name:   getEnvString("APP_NAME", "AuthBase"),
	}
	getEnvJWTSecretSpecific()
}

func getEnvString(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	if fallback == "" {
		log.Fatalf("%s cant be empty. Please fill out in .env", key)
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		return v == "true" || v == "1"
	}
	return fallback
}

func getEnvJWTSecretSpecific() {
	if v := os.Getenv("JWT_SECRET"); v != "" {
		JwtSecret = []byte(v)
	} else {
		log.Fatalf("JWT cant be empty. Please fill out in .env")

	}
}
