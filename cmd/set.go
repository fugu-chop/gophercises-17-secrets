package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an encrypted key",
	Long: `Use the 'set' command, along with a key name, a
	key value, followed by the -k flag with an encoding key
	to store an encrypted key locally.`,
	// Rely on cobra PositionalArgs for 'non-named' flags
	Args: cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Println(arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.
	setCmd.Flags().StringP("encryptionKey", "k", "", `The keyName flag (-k) is 
	used to provide an encryption key to encrypt a key`)
	setCmd.MarkFlagRequired("encryptionKey")
}

func encryptKey(value, encryptionKey string) error {
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

	return nil
}
