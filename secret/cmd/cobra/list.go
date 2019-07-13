package cobra

import (
	"fmt"
	"secret"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your keys in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		value, err := v.List()
		if err != nil {
			fmt.Println(`Something went wrong while listing your secret. Double check if you included the right encoding key.`)
			return
		}
		for _, v := range value {
			fmt.Println(v)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
