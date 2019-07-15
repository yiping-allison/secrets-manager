package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yiping-allison/secrets-manager/secret"
)

var acsvCmd = &cobra.Command{
	Use:   "appendcsv",
	Short: "Appends data from a csv file to an existing secrets file",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		v := secret.File(encodingKey, secretsPath())
		err := v.AppendCSV(filename)
		if err != nil {
			fmt.Println("An error occurred when trying to append your csv values. Check if you included your correct encoding key or filename.")
			return
		}
		fmt.Println("Successfully appended your csv file!")
	},
}

func init() {
	RootCmd.AddCommand(acsvCmd)
}
