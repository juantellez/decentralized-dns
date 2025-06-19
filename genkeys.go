package main

import (
	"fmt"
	"log"

	dnsclient "github.com/juantellez/dns-chain/dns_client"
)

func main() {
	pub, priv, err := dnsclient.GenerateKeyPair()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Public Key:", pub)
	fmt.Println("Private Key:", priv)
}
