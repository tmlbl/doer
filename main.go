package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var configPath string
var certPath string
var privKeyPath string
var port int

func main() {
	flag.StringVar(&configPath, "config", "doer.json", "Path to a config file")
	flag.StringVar(&certPath, "cert", "", "Path to a certificate file")
	flag.StringVar(&privKeyPath, "key", "", "Path to a private key")
	flag.IntVar(&port, "port", 8778, "Port to listen on")
	flag.Parse()

	config := loadConfig(configPath)
	server := Server{config}

	addr := fmt.Sprintf(":%d")
	if certPath == "" {
		log.Println("Starting the server with plain HTTP on port", port)
		http.ListenAndServe(":8778", &server)
	} else {
		log.Println("Starting the server with HTTPS on port", port)
		http.ListenAndServeTLS(":8778", certPath, privKeyPath, &server)
	}
}
