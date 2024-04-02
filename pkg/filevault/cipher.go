package filevault

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
	FilePath string
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
