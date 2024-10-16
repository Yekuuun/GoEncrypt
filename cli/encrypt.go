package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdDir *cobra.Command

func runEncryption(cmd *cobra.Command, args []string) error {
	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return fmt.Errorf("[!] ERROR : no file path founded")
	}

	fmt.Println(filePath)

	return nil
}

func init() {
	CmdDir = &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a given file using it's path",
		RunE:  runEncryption,
	}

	CmdDir.Flags().StringP("file", "f", "", "Path to the file to encrypt")
	CmdDir.MarkFlagRequired("file")

	rootCmd.AddCommand(CmdDir)
}
