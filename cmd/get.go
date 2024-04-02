package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch an already encrypted key",
	Long: `Use the 'get' command, along with a key name followed
	by the -d flag with an encoding key to retrieve an encrypted 
	key stored locally.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Println(arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.
	getCmd.Flags().StringP("decryptionKey", "k", "", `The decryption flag (-k) is
	// used to provide an decryption key to decrypt a stored key`)
	getCmd.MarkFlagRequired("decryptionKey")
}
