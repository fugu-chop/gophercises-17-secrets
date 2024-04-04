package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Initialise this on startup
// The CLI doesn't seem to allow persistent running app
// so memory will be wiped on every go run
type FileVault struct {
	// Where to write the file
	FilePath      string
	EncryptionKey string
}

// Rely on position args when calling
func (f *FileVault) Set(value, encryptionKey string) error {
	key, _ := hex.DecodeString(encryptionKey)
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
	// Need to figure out how to write to file
	return nil
}

func (f *FileVault) Get(value, encryptionKey string) error {
	key, _ := hex.DecodeString(encryptionKey)
	// Pull this value from db
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
