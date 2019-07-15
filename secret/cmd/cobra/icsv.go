package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yiping-allison/secrets-manager/secret"
)

var icsvCmd = &cobra.Command{
	Use:   "importcsv",
	Short: "Imports a csv file to a new secrets file",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		err := secret.ImportCSV(encodingKey, secretsPath(), filename)
		if err != nil {
			fmt.Println("An error occurred while trying to import your csv file.",
				"Please make sure you have the right encoding key, filepath, and you're starting with no existing",
				"secrets file.")
			return
		}
		fmt.Println("Successfully imported csv file!")
	},
}

func init() {
	RootCmd.AddCommand(icsvCmd)
}
