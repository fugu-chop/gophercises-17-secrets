package main

import (
	"fmt"
	"log"
	"os"
	vault "secrets/pkg/vault"
)

const (
	secretsLocation = "/Users/dean/Desktop/secrets.txt"
	encryptionKey   = "6368616e676520746869732070617373"
)

func main() {
	// Temporarily disable CLI
	// cmd.Execute()
	vault := vault.FileVault{
		EncryptionKey: encryptionKey,
	}

	// Check if secrets files exists
	if _, err := os.Stat(secretsLocation); err != nil {
		log.Printf("creating secrets.txt file at %s", secretsLocation)
		vault.VaultFile, err = os.Create(secretsLocation)
		if err != nil {
			log.Fatalf("could not create secrets file: %s", err)
		}
	} else {
		vault.VaultFile, err = os.Open(secretsLocation)
		if err != nil {
			log.Fatalf("could not open secrets file: %s", err)
		}
		log.Println("secrets.txt file already exists")
	}

	fmt.Println(vault.VaultFile.Name())
}
