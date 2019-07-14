package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yiping-allison/secrets-manager/secret"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key := args[0]
		err := v.Remove(key)
		if err != nil {
			fmt.Println(`Something went wrong while removing your secret. Double check if you included the right encoding key.`)
			return
		}
		fmt.Println("Secret removed succesfully!")
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
