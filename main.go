package main

import (
	"fmt"
	"log"
	"os"
	vault "secrets/pkg/vault"
)

func main() {
	// Temporarily disable CLI
	// cmd.Execute()
	vault := vault.FileVault{
		EncryptionKey: "6368616e676520746869732070617373",
		FilePath:      "/Users/dean/Desktop",
	}

	var secretsFile *os.File
	// Check if secrets files exists
	if _, err := os.Stat(vault.FilePath + "/secrets.txt"); err != nil {
		log.Printf("creating secrets.txt file at %s", vault.FilePath)
		secretsFile, err = os.Create(vault.FilePath + "/secrets.txt")
		if err != nil {
			log.Fatalf("could not create secrets file: %s", err)
		}
	} else {
		secretsFile, err = os.Open(vault.FilePath + "/secrets.txt")
		if err != nil {
			log.Fatalf("could not open secrets file: %s", err)
		}
	}

	fmt.Println(secretsFile.Name())
}
