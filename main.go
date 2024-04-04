package main

import (
	"fmt"
	"log"
	"os"
	vault "secrets/pkg/vault"
)

const (
	secretsLocation = "/Users/dean/Desktop"
	encryptionKey   = "6368616e676520746869732070617373"
)

func main() {
	// Temporarily disable CLI
	// cmd.Execute()
	vault := vault.FileVault{
		EncryptionKey: encryptionKey,
	}

	// Shut the linter up
	_ = vault

	var secretsFile *os.File
	// Check if secrets files exists
	if _, err := os.Stat(secretsLocation); err != nil {
		log.Printf("creating secrets.txt file at %s", secretsLocation)
		secretsFile, err = os.Create(secretsLocation)
		if err != nil {
			log.Fatalf("could not create secrets file: %s", err)
		}
	} else {
		secretsFile, err = os.Open(secretsLocation)
		if err != nil {
			log.Fatalf("could not open secrets file: %s", err)
		}
		log.Println("secrets.txt file already exists")
	}

	// Write the file to a map, insert onto filevault
	fmt.Println(secretsFile.Name())
}
