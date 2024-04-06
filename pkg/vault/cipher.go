package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

type FileVault struct {
	EncryptionKey string
	vaultSecrets  map[string]string
}

func (f *FileVault) GenerateVault(fileLocation string) error {
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		return err
	}

	if len(file) == 0 {
		return nil
	}

	f.vaultSecrets = make(map[string]string)
	secretsPairs := strings.Split(string(file), "\n")
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

// Ideally we only write secrets that are new to disk
// But not a hard requirement
func (f *FileVault) WriteSecrets(secrets map[string]string, writePath string) error {
	return nil
}

// Does not allow for amending existing
func (f *FileVault) Set(value string) error {
	key, _ := hex.DecodeString(f.EncryptionKey)
	plaintext := []byte(value)
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

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// ciphertext is now encrypted
	fmt.Println(ciphertext)
	// add to map
	// write to disc
	return nil
}

func (f *FileVault) Get(value string) error {
	key, _ := hex.DecodeString(f.EncryptionKey)
	// Pull this value from file
	ciphertext, _ := hex.DecodeString("")

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %s", err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return fmt.Errorf("ciphertext is too short: %s", err)
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("%s", ciphertext)

	return nil
}
