package cmd

import (
	"fmt"
	"log"
	vault "secrets/pkg/vault"

	"github.com/spf13/cobra"
)

const (
	secretsLocation = "/Users/dean/Desktop/secrets.txt"
	encryptionKey   = "6368616e676520746869732070617373"
)

var decryptionKey string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch an already encrypted key",
	Long: `Use the 'get' command, along with a key name followed
	by the -k flag with an encoding key to retrieve an encrypted 
	key stored locally.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		vault := vault.FileVault{
			EncryptionKey: encryptionKey,
		}

		if err := vault.GenerateVault(secretsLocation); err != nil {
			log.Fatalf("could not generate vault from secrets file: %s", err)
		}

		if decryptionKey == vault.EncryptionKey {
			key, err := vault.Get(args[0])
			if err != nil {
				log.Fatalf("could not fetch key %s, %s", args[0], err)
			}
			fmt.Println(key)
		} else {
			log.Fatal("encryption key does not match")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.
	getCmd.Flags().StringVarP(&decryptionKey, "decryptionKey", "k", "", `The decryption flag (-k) is
	// used to provide an decryption key to decrypt a stored key`)
	getCmd.MarkFlagRequired("decryptionKey")
}
