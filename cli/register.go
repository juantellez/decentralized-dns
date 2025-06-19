package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	dnsclient "github.com/juantellez/dns-chain/dns_client"
)

const dataFile = "dns_client/domains.json"

func normalizeDomain(domain string) string {
	if !strings.HasSuffix(domain, ".") {
		return domain + "."
	}
	return domain
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run cli/register.go <domain> <owner-public-key-base64> <ns1> [<ns2> ...]")
		return
	}

	domain := normalizeDomain(os.Args[1])
	owner := os.Args[2]
	ns := os.Args[3:]

	data := make(map[string]dnsclient.DomainRecord)

	content, err := os.ReadFile(dataFile)
	if err == nil {
		_ = json.Unmarshal(content, &data)
	}

	if _, exists := data[domain]; exists {
		fmt.Printf("[ERROR] Domain %s already exists.\n", domain)
		return
	}

	data[domain] = dnsclient.DomainRecord{
		Owner:      owner,
		NS:         ns,
		Expiration: time.Now().Add(365 * 24 * time.Hour).Unix(),
	}

	updated, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile(dataFile, updated, 0644)

	fmt.Printf("[SUCCESS] Registered %s with owner %s and NS records: %v\n", domain, owner, ns)
}
