package config

import "flag"

var FlagServerAddress string
var FlagBaseURL string
var FlagLogLevel string

const defaultServerAddress = "localhost:8080"
const defaultBaseURL = "http://localhost:8080"
const defaultLogLevel = "info"

func ParseFlags() {
	flag.StringVar(&FlagServerAddress, "a", defaultServerAddress, "Server address")
	flag.StringVar(&FlagBaseURL, "b", defaultBaseURL, "Base url before shortened url")
	flag.StringVar(&FlagLogLevel, "l", defaultLogLevel, "Logging level")
	flag.Parse()

	if FlagServerAddress == "" {
		FlagServerAddress = defaultServerAddress
	}
	if FlagBaseURL == "" {
		FlagBaseURL = defaultBaseURL
	}
	if FlagLogLevel == "" {
		FlagLogLevel = defaultLogLevel
	}
}
