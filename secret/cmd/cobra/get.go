package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yiping-allison/secrets-manager/secret"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println(`Something went wrong while retrieving your secret. Double check if you included the right encoding key.`)
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
