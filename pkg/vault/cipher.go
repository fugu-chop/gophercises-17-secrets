package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
Encryption and decryption are based on crypto/cipher
package available as part of standard library
See: https://pkg.go.dev/crypto/cipher
*/
type FileVault struct {
	EncryptionKey string
	vaultSecrets  map[string]string
	fileLocation  string
}

func (f *FileVault) GenerateVault(fileLocation string) error {
	f.fileLocation = fileLocation

	// Ensure file exists through O_APPEND or O_CREATE
	file, err := os.OpenFile(f.fileLocation, os.O_APPEND|os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("could not access secrets file: %s", err)
	}
	defer file.Close()

	buffer := make([]byte, 256)
	// Reuse buffer
	for {
		_, err := file.Read(buffer)
		// io.EOF defines the end of the file
		// but is returned as an error
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return err
		}
	}

	f.vaultSecrets = make(map[string]string)
	secretsPairs := strings.Split(string(buffer), "\n")
	for _, secret := range secretsPairs {
		pair := strings.Split(secret, " ")
		// This len check is to ensure that the last entry of the
		// file which will be a blank line is not parsed into
		// vaultSecrets so that an out-of-bounds error won't occur
		if len(pair) > 1 {
			f.vaultSecrets[pair[0]] = pair[1]
		}
	}

	return nil
}

func (f *FileVault) WriteSecrets(secrets map[string]string) error {
	file, err := os.OpenFile(f.fileLocation, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	for key, secret := range secrets {
		file.WriteString(fmt.Sprintf("%s %s\n", key, secret))
	}

	return nil
}

func (f *FileVault) Set(flag, secret string) error {
	if _, ok := f.vaultSecrets[flag]; ok {
		return nil
	}

	key, _ := hex.DecodeString(f.EncryptionKey)
	plaintext := []byte(secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %s", err)
	}

	// Generate a unique (but not necessarily secure) IV
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return fmt.Errorf("failed to generate ciphertext: %s", err)
	}

	// ciphertext is mutated
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	// We need to format the bytes to "base 16, lower-case, two characters per byte"
	f.vaultSecrets[flag] = fmt.Sprintf("%x", ciphertext)

	// we will have to replace the file on each iteration
	if err = f.WriteSecrets(f.vaultSecrets); err != nil {
		return fmt.Errorf("failed to write secrets: %s", err)
	}

	return nil
}

func (f *FileVault) Get(value string) (string, error) {
	if _, ok := f.vaultSecrets[value]; !ok {
		return "", fmt.Errorf("failed find key in vault: %s", value)
	}

	key, _ := hex.DecodeString(f.EncryptionKey)
	ciphertext, err := hex.DecodeString(f.vaultSecrets[value])
	if err != nil {
		return "", fmt.Errorf("failed to decrypt cipher: %s", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %s", err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext is too short: %s", err)
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%q", ciphertext), nil
}
