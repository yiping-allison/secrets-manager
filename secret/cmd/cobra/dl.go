package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yiping-allison/secrets-manager/secret"
)

var dlCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an existing secret file from your user directory",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		err := v.Delete()
		if err != nil {
			fmt.Println(`Something went wrong while deleting your existing Secrets file. Double check if you included the right encoding key.`)
			return
		}
		fmt.Println("Secret file removed succesfully!")
	},
}

func init() {
	RootCmd.AddCommand(dlCmd)
}
