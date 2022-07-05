package cmd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/spf13/cobra"
	"github.com/viostream/fusion-theme-sync/theme"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ensureAPIKeyPresent()
		fmt.Println("Running fusion-theme-sync put")

		// Load substitution keys
		var substitutions map[string]string
		if substitutionFile != "" {
			var err error
			substitutions, err = theme.LoadSubstitutionsFromDisk(substitutionFile)
			if err != nil {
				log.Fatal(err)
			}
		}

		// Load theme from disk
		theme, err := theme.LoadFromDisk(directory, substitutions)
		if err != nil {
			log.Fatal(err)
		}

		url, _ := url.Parse(fusionURL)
		c := fusionauth.NewClient(&http.Client{}, url, apiKey)

		fmt.Printf("Writing theme to FusionAuth at %v...\n", fusionURL)

		// The SDK doesn't really contain a proper way of patching; we have to build our own
		// request body
		patchThemeRequest := map[string]interface{}{
			"theme": map[string]interface{}{
				"defaultMessages": theme.DefaultMessages,
				"stylesheet":      theme.Stylesheet,
				"templates":       theme.Templates,
			},
		}

		_, ferr, err := c.PatchTheme(themeID, patchThemeRequest)
		if err != nil {
			log.Fatal(err)
		}

		if ferr != nil {
			log.Printf("[ERROR] %v", ferr)
			log.Fatal("[ERROR] Got some error from FusionAuth but the SDK isn't great so not sure what - maybe check your API key")
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
