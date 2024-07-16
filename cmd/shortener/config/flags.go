package config

import "flag"

var AppRunServerAddressFromFlag string
var RedirectAddressFromFlag string

func ParseFlags() {
	flag.StringVar(&AppRunServerAddressFromFlag, "a", "http://localhost:8080", "Address to listen for incoming requests")
	flag.StringVar(&RedirectAddressFromFlag, "b", "http://localhost:8081", "Address to redirect")
	flag.Parse()
}
