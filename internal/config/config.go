package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

var serviceVersion = "local"

const (
	port = "HTTP_PORT"

	postgresHost    = "POSTGRES_HOST"
	postgresUser    = "POSTGRES_USER"
	postgresPass    = "POSTGRES_PASSWORD"
	postgresDB      = "POSTGRES_DATABASE"
	postgresTimeOUT = "POSTGRES_TIMEOUT"

	freeCurrencyApi    = "FREE_CURRENCY_URL"
	freeCurrencyApiKey = "FREE_CURRENCY_API_KEY"
	freeCurrencyTime   = "FREE_CURRENCY_TIME"

	timeoutSeconds = "TIMEOUT_SECONDS"
)

type Config struct {
	Port               string
	DbPostgresUrl      string
	DbTimeOUT          int
	FreeCurrencyApi    string
	FreeCurrencyApiKey string
	FreeCurrencyTime   int
	TimeoutSeconds     int
}

func New() Config {
	postgresURL := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv(postgresUser),
		os.Getenv(postgresPass),
		os.Getenv(postgresHost),
		os.Getenv(postgresDB))

	return Config{
		Port:               GetEnvString(port, ""),
		DbPostgresUrl:      postgresURL,
		DbTimeOUT:          GetEnvFloat(postgresTimeOUT, 15),
		FreeCurrencyApi:    GetEnvString(freeCurrencyApi, ""),
		FreeCurrencyApiKey: GetEnvString(freeCurrencyApiKey, ""),
		FreeCurrencyTime:   GetEnvFloat(freeCurrencyTime, 1),
		TimeoutSeconds:     GetEnvFloat(timeoutSeconds, 60),
	}
}

func GetVersion() string {
	return serviceVersion
}

func GetEnvString(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultValue
}

func GetEnvFloat(key string, def int) int {
	val, err := strconv.Atoi(GetEnvString(key, fmt.Sprintf("%d", def)))
	if err != nil {
		return def
	}
	return val
}
