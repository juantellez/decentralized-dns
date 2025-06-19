package main

import (
	"bufio"
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter private key (base64): ")
	privBase64, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read private key:", err)
	}
	privBase64 = strings.TrimSpace(privBase64)

	privBytes, err := base64.StdEncoding.DecodeString(privBase64)
	if err != nil {
		log.Fatal("Invalid base64 private key:", err)
	}

	fmt.Print("Enter message to sign: ")
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read message:", err)
	}
	message = strings.TrimSpace(message)

	signature := ed25519.Sign(privBytes, []byte(message))
	signBase64 := base64.StdEncoding.EncodeToString(signature)

	fmt.Println("Signature (base64):", signBase64)
}
