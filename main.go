package main

import (
	"log"

	"github.com/juantellez/dns-chain/config"
	dnsclient "github.com/juantellez/dns-chain/dns_client"
	"github.com/miekg/dns"
)

func main() {
	config.LoadConfig("config.json")

	dns.HandleFunc(".", dnsclient.HandleDNSRequest) // usa el prefijo del paquete
	server := &dns.Server{Addr: ":53", Net: "udp"}
	log.Println("[INFO] DNS server listening on UDP port 53")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("[FATAL] Failed to start server: %v", err)
	}
}
