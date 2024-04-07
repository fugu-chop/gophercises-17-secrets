package main

import (
	"log"
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

	if err := vault.GenerateVault(secretsLocation); err != nil {
		log.Fatalf("could not generate vault from secrets file: %s", err)
	}

	// Test what writing a map[string]string to a file looks like
	// probably need to use encoding/gob
}
