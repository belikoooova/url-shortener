package config

import "flag"

var AppRunServerAddressFromFlag string
var RedirectAddressFromFlag string

const defaultAddr = "http://localhost:8080"

func ParseFlags() {
	flag.StringVar(&AppRunServerAddressFromFlag, "a", defaultAddr, "Address to listen for incoming requests")
	flag.StringVar(&RedirectAddressFromFlag, "b", defaultAddr, "Address to redirect")
	flag.Parse()

	if AppRunServerAddressFromFlag == "" {
		AppRunServerAddressFromFlag = defaultAddr
	}
	if RedirectAddressFromFlag == "" {
		RedirectAddressFromFlag = defaultAddr
	}
}
