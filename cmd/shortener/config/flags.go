package config

import "flag"

var AppRunServerAddressFromFlag string
var RedirectAddressFromFlag string

func ParseFlags() {
	flag.StringVar(&AppRunServerAddressFromFlag, "a", "localhost:8080", "Address to listen for incoming requests")
	flag.StringVar(&RedirectAddressFromFlag, "b", "http://localhost:8080", "Address to redirect")
	flag.Parse()

	if AppRunServerAddressFromFlag == "" {
		AppRunServerAddressFromFlag = "localhost:8080"
	}
	if RedirectAddressFromFlag == "" {
		RedirectAddressFromFlag = "http://localhost:8080"
	}
}
