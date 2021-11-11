package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	apiKey           string
	directory        string
	fusionURL        string
	substitutionFile string
	themeID          string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fusion-theme-sync",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&directory, "dir", "", "The local dir to store/retrieve files")
	rootCmd.MarkPersistentFlagRequired("dir")

	rootCmd.PersistentFlags().StringVar(&fusionURL, "fusion-url", "", "The base URL of the fusionauth server")
	rootCmd.MarkPersistentFlagRequired("fusion-url")

	rootCmd.PersistentFlags().StringVar(&substitutionFile, "substitution-file", "", "Path to a file containing default messages to substitute")

	rootCmd.PersistentFlags().StringVar(&themeID, "theme-id", "", "The ID of the fusionauth theme")
	rootCmd.MarkPersistentFlagRequired("theme-id")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func ensureAPIKeyPresent() {
	// API key must be set as an ENV or typed in - we don't accept it via a flag
	// for security reasons (command history etc)
	if apiKey != "" {
		return
	}

	// check ENVs first
	var ok bool
	apiKey, ok = os.LookupEnv("FA_API_KEY")

	// Otherwise ask for input
	if !ok {
		fmt.Printf("Please enter your FusionAuth API key:\n > ")
		fmt.Scanln(&apiKey)
	}
}
