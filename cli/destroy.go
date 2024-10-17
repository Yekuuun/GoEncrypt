/*
* Notes : destroying keys stored in app.
 */

package cli

import (
	"GoEncrypt/pkg/keys"
	"fmt"

	"github.com/spf13/cobra"
)

var CmdDestroy *cobra.Command

func RunDestroy(cmd *cobra.Command, args []string) error {
	fmt.Printf("---------------------")
	fmt.Println("\nDestorying keys.....")

	err := keys.CleanKeys()
	if err != nil {
		return err
	}

	fmt.Println("Keys deleted........")

	fmt.Printf("\n Bye ;)")
	return nil
}

func init() {
	CmdDestroy = &cobra.Command{
		Use:   "destroy",
		Short: "Destroy all keys stored in /data/keys",
		RunE:  RunDestroy,
	}

	rootCmd.AddCommand(CmdDestroy)
}
