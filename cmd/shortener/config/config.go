package config

import "os"

type Config struct {
	ServerAddress string
	BaseUrl       string
}

func Configure() *Config {
	ParseFlags()
	serverAddress := getEnvOrDefault("SERVER_ADDRESS", ServerAddressFromFlag)
	baseUrl := getEnvOrDefault("BASE_URL", BaseUrlFromFlag)
	return &Config{
		ServerAddress: serverAddress,
		BaseUrl:       baseUrl,
	}
}

func getEnvOrDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); value != "" && ok {
		return value
	}
	return fallback
}
