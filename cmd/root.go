package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "secrets",
	Short: "A tool to store and fetch encrypted files",
	Long: `This is a package that provides a CLI tool to store and
	retrieve encrypted files, much like Hashicorp's Vault product.
	You will need to use either:
	1. The 'set' command with a key name, followed by the -k flag
		and an encoding key to set an encrypted value; or
	2. The 'get' command with a key name, followed by the -k flag
		with your encoding key to fetch an encrypted value.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
