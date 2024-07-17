package config

import "flag"

var ServerAddressFromFlag string
var BaseUrlFromFlag string

const defaultServerAddress = "localhost:8080"
const defaultBaseUrl = "http://localhost:8080"

func ParseFlags() {
	flag.StringVar(&ServerAddressFromFlag, "a", defaultServerAddress, "Server address")
	flag.StringVar(&BaseUrlFromFlag, "b", defaultBaseUrl, "Base url before shortened url")
	flag.Parse()

	if ServerAddressFromFlag == "" {
		ServerAddressFromFlag = defaultServerAddress
	}
	if BaseUrlFromFlag == "" {
		BaseUrlFromFlag = defaultBaseUrl
	}
}
