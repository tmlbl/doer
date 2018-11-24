package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Tasks []Task `json:"tasks"`
}

func loadConfig(fpath string) *Config {
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatalln(err)
	}
	cfg := Config{}
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return &cfg
}
