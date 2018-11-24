package main

import (
	"flag"
	"net/http"
)

var configPath string
var certPath string
var privKeyPath string

func main() {
	flag.StringVar(&configPath, "config", "doer.json", "Path to a config file")
	flag.StringVar(&certPath, "cert", "", "Path to a certificate file")
	flag.StringVar(&privKeyPath, "key", "", "Path to a private key")
	flag.Parse()

	config := loadConfig(configPath)
	server := Server{config}

	if certPath == "" {
		http.ListenAndServe(":8778", &server)
	} else {
		http.ListenAndServeTLS(":8778", certPath, privKeyPath, &server)
	}
}
