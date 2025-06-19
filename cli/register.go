package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

const dataFile = "dns_client/domains.json"

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run cli/register.go <domain> <owner-public-key-base64> <ns1> [<ns2> ...]")
		return
	}

	domain := os.Args[1]
	owner := os.Args[2]
	ns := os.Args[3:]

	data := make(map[string]DomainRecord)

	content, err := os.ReadFile(dataFile)
	if err == nil {
		_ = json.Unmarshal(content, &data)
	}

	if _, exists := data[domain]; exists {
		fmt.Printf("[ERROR] Domain %s already exists.\n", domain)
		return
	}

	data[domain] = DomainRecord{
		Owner: owner,
		NS:    ns,
	}

	updated, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile(dataFile, updated, 0644)

	fmt.Printf("[SUCCESS] Registered %s with owner %s and NS records: %v\n", domain, owner, ns)
}
