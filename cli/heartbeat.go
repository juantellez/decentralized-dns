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

	record, exists := data[domain]
	if !exists {
		fmt.Printf("[ERROR] Domain %s not found.\n", domain)
		return
	}

	// Extiende el tiempo de expiración por un año
	record.Expiration = time.Now().Add(365 * 24 * time.Hour).Unix()
	data[domain] = record

	updated, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("[ERROR] Failed to serialize updated data:", err)
		return
	}

	err = os.WriteFile(dataFile, updated, 0644)
	if err != nil {
		fmt.Println("[ERROR] Could not write to domains.json:", err)
		return
	}

	fmt.Printf("[HEARTBEAT] Domain %s extended until %s\n", domain, time.Unix(record.Expiration, 0).UTC().Format(time.RFC3339))
}
