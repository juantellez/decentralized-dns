package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	FallbackDNSServers []string `json:"fallback_dns_servers"`
}

var Conf Config

func LoadConfig(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("[FATAL] Failed to read config file: %v", err)
	}
	err = json.Unmarshal(file, &Conf)
	if err != nil {
		log.Fatalf("[FATAL] Failed to parse config file: %v", err)
	}
}
