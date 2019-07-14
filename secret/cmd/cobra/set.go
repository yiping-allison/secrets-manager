package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yiping-allison/secrets-manager/secret"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			fmt.Println(`Something went wrong while setting your secret. Double check if you included the right key.`)
			return
		}
		fmt.Println("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
