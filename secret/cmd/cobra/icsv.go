package cobra

import (
	"github.com/spf13/cobra"
	"github.com/yiping-allison/secrets-manager/secret"
)

var icsvCmd = &cobra.Command{
	Use:   "importcsv",
	Short: "Imports a csv file to a new secrets file",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		secret.ImportCSV(encodingKey, secretsPath(), filename)
	},
}

func init() {
	RootCmd.AddCommand(icsvCmd)
}
