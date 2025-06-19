package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	dnsclient "github.com/juantellez/dns-chain/dns_client"
)

const dataFile = "dns_client/domains.json"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run cli/heartbeat.go <domain>")
		return
	}

	domain := os.Args[1]
	data := make(map[string]dnsclient.DomainRecord)

	content, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("[ERROR] Could not read domains.json:", err)
		return
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("[ERROR] Invalid JSON format:", err)
		return
	}

	if _, exists := data[domain]; !exists {
		fmt.Printf("[ERROR] Domain %s not found.\n", domain)
		return
	}

	// En el futuro: actualizar timestamp en la estructura
	fmt.Printf("[HEARTBEAT] Domain %s is alive at %s\n", domain, time.Now().UTC().Format(time.RFC3339))
}
