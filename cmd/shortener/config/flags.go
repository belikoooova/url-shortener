package config

import "flag"

var AppRunServerAddressFromFlag string
var RedirectAddressFromFlag string

func ParseFlags() {
	flag.StringVar(&AppRunServerAddressFromFlag, "a", ":8080", "Address to listen for incoming requests")
	flag.StringVar(&RedirectAddressFromFlag, "b", ":8080", "Address to redirect")
	flag.Parse()

	if AppRunServerAddressFromFlag == "" {
		AppRunServerAddressFromFlag = ":8080"
	}
	if RedirectAddressFromFlag == "" {
		RedirectAddressFromFlag = ":8080"
	}
}
