package config

import "os"

type Config struct {
	AppRunServerAddress   string
	RedirectServerAddress string
}

func Configure() *Config {
	ParseFlags()
	appRunServerAddress := getEnvOrDefault("SERVER_ADDRESS", AppRunServerAddressFromFlag)
	redirectAddress := getEnvOrDefault("BASE_URL", RedirectAddressFromFlag)
	return &Config{
		AppRunServerAddress:   appRunServerAddress,
		RedirectServerAddress: redirectAddress,
	}
}

func getEnvOrDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); value != "" && ok {
		return value
	}
	return fallback
}
