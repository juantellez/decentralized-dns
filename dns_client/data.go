package dnsclient

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type DomainRecord struct {
	Owner string   `json:"owner"` // base64-encoded public key
	NS    []string `json:"ns"`
}

var blockchainRecords = make(map[string]DomainRecord)

func init() {
	data, err := os.ReadFile("dns_client/domains.json")
	if err != nil {
		log.Fatalf("[FATAL] Could not read domains.json: %v", err)
	}

	err = json.Unmarshal(data, &blockchainRecords)
	if err != nil {
		log.Fatalf("[FATAL] Could not parse domains.json: %v", err)
	}

	log.Printf("[INFO] Loaded %d domain records from domains.json", len(blockchainRecords))
}

func domainInBlockchain(domain string) bool {
	_, exists := blockchainRecords[domain]
	return exists
}

func getTLD(domain string) string {
	parts := strings.Split(strings.TrimSuffix(domain, "."), ".")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
