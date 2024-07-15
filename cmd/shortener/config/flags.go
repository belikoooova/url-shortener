package config

import "flag"

var AppRunAddress string
var RedirectAddress string

func ParseFlags() {
	flag.StringVar(&AppRunAddress, "a", "localhost:8080", "Address to listen for incoming requests")
	flag.StringVar(&RedirectAddress, "b", "localhost:8080", "Address to redirect")
	flag.Parse()
}
