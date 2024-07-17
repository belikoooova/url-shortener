package config

import "os"

type Config struct {
	ServerAddress string
	BaseURL       string
	LogLevel      string
}

func Configure() *Config {
	ParseFlags()
	serverAddress := getEnvOrDefault("SERVER_ADDRESS", FlagServerAddress)
	baseURL := getEnvOrDefault("BASE_URL", FlagBaseURL)
	logLevel := getEnvOrDefault("LOG_LEVEL", FlagLogLevel)
	return &Config{
		ServerAddress: serverAddress,
		BaseURL:       baseURL,
		LogLevel:      logLevel,
	}
}

func getEnvOrDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); value != "" && ok {
		return value
	}
	return fallback
}
