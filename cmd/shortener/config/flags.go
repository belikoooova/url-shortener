package config

import "flag"

var ServerAddressFromFlag string
var BaseURLFromFlag string

const defaultServerAddress = "localhost:8080"
const defaultBaseURL = "http://localhost:8080"

func ParseFlags() {
	flag.StringVar(&ServerAddressFromFlag, "a", defaultServerAddress, "Server address")
	flag.StringVar(&BaseURLFromFlag, "b", defaultBaseURL, "Base url before shortened url")
	flag.Parse()

	if ServerAddressFromFlag == "" {
		ServerAddressFromFlag = defaultServerAddress
	}
	if BaseURLFromFlag == "" {
		BaseURLFromFlag = defaultBaseURL
	}
}
