package config

import "flag"

var AppRunServerAddressFromFlag string
var RedirectAddressFromFlag string

func ParseFlags() {
	flag.StringVar(&AppRunServerAddressFromFlag, "a", "localhost:8081", "Address to listen for incoming requests")
	flag.StringVar(&RedirectAddressFromFlag, "b", "localhost:8080", "Address to redirect")
	flag.Parse()
}
