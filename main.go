package main

import (
	"log"
	"os"
	vault "secrets/pkg/vault"
)

const (
	secretsLocation = "/Users/dean/Desktop/secrets.txt"
	encryptionKey   = "6368616e6765207468697320ab70617373"
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
		file, err := os.Create(secretsLocation)
		if err != nil {
			log.Fatalf("could not create secrets file: %s", err)
		}
		defer file.Close()
	} else {
		log.Print("secrets.txt file already exists")
	}

	if err := vault.GenerateVault(secretsLocation); err != nil {
		log.Fatalf("could not read from secrets file: %s", err)
	}

	// Test what writing a map[string]string to a file looks like
	// probably need to use encoding/gob
}
