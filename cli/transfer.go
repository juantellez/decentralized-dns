package main

import (
	"encoding/json"
	"fmt"
	"os"

	dnsclient "github.com/juantellez/dns-chain/dns_client"
)

const dataFile = "dns_client/domains.json"

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: go run cli/transfer.go <domain> <new-owner-public-key-base64> <signature-base64> <message>")
		return
	}

	domain := os.Args[1]
	newOwner := os.Args[2]
	signature := os.Args[3]
	message := []byte(os.Args[4])

	data := make(map[string]dnsclient.DomainRecord)
	content, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("[ERROR] Could not read domains.json:", err)
		return
	}
	_ = json.Unmarshal(content, &data)

	rec, exists := data[domain]
	if !exists {
		fmt.Printf("[ERROR] Domain %s not found.\n", domain)
		return
	}

	// Verificar firma del mensaje con la clave pública actual
	valid, err := dnsclient.VerifySignature(rec.Owner, message, signature)
	if err != nil {
		fmt.Println("[ERROR] Failed to verify signature:", err)
		return
	}
	if !valid {
		fmt.Println("[ERROR] Invalid signature.")
		return
	}

	// Transferir propiedad
	rec.Owner = newOwner
	data[domain] = rec

	updated, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile(dataFile, updated, 0644)

	fmt.Printf("[SUCCESS] Domain %s transferred to new owner %s\n", domain, newOwner)
}
