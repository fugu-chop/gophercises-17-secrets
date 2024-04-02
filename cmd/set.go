/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
