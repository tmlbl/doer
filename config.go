package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Tasks []Task `json:"tasks"`
}

// JSON representation without the secrets
func (c Config) toCleanJSON() []byte {
	for i := range c.Tasks {
		c.Tasks[i].Secret = ""
	}
	data, _ := json.Marshal(&c)
	return data
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
