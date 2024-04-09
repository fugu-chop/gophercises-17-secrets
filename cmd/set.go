package cmd

import (
	"log"
	vault "secrets/pkg/vault"

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
	// Need to validate second arg has quotation marks (or no quotation marks?)
	Args: cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		vault := vault.FileVault{
			EncryptionKey: encryptionKey,
		}

		if err := vault.GenerateVault(secretsLocation); err != nil {
			log.Fatalf("could not generate vault from secrets file: %s", err)
		}

		err := vault.Set(args[0], args[1])
		if err != nil {
			log.Fatal(err)
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
