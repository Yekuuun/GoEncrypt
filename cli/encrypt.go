package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdDir *cobra.Command

func runEncryption(cmd *cobra.Command, arg []string) error {
	fmt.Println("running command")

	return nil
}

func init() {
	CmdDir = &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a given file using it's path",
		RunE:  runEncryption,
	}

	rootCmd.AddCommand(CmdDir)
}
