package cli

import (
	"encoding/json"
	"fmt"
	"os"
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

	data := make(map[string]DomainRecord)
	content, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("[ERROR] Could not read domains.json")
		return
	}
	_ = json.Unmarshal(content, &data)

	rec, exists := data[domain]
	if !exists {
		fmt.Printf("[ERROR] Domain %s not found.\n", domain)
		return
	}

	// Verificar firma con owner actual
	valid, err := VerifySignature(rec.Owner, message, signature)
	if err != nil || !valid {
		fmt.Println("[ERROR] Invalid signature")
		return
	}

	// Actualizar propietario
	rec.Owner = newOwner
	data[domain] = rec

	updated, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile(dataFile, updated, 0644)

	fmt.Printf("[SUCCESS] Domain %s transferred to new owner %s\n", domain, newOwner)
}
