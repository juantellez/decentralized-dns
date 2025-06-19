package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run cli/heartbeat.go <domain>")
		return
	}

	domain := os.Args[1]

	data := make(map[string][]string)
	content, err := os.ReadFile("dns_client/domains.json")
	if err != nil {
		fmt.Println("[ERROR] Could not read domains.json")
		return
	}
	_ = json.Unmarshal(content, &data)

	if _, exists := data[domain]; !exists {
		fmt.Printf("[ERROR] Domain %s not found.\n", domain)
		return
	}

	// MVP: solo imprimimos, en el futuro se guardaría el timestamp
	fmt.Printf("[HEARTBEAT] Domain %s is alive at %s\n", domain, time.Now().UTC().Format(time.RFC3339))
}
