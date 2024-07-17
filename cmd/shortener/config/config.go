package config

import "os"

type Config struct {
	ServerAddress string
	BaseURL       string
}

func Configure() *Config {
	ParseFlags()
	serverAddress := getEnvOrDefault("SERVER_ADDRESS", ServerAddressFromFlag)
	baseUrl := getEnvOrDefault("BASE_URL", BaseURLFromFlag)
	return &Config{
		ServerAddress: serverAddress,
		BaseURL:       baseUrl,
	}
}

func getEnvOrDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); value != "" && ok {
		return value
	}
	return fallback
}
