package cli

import (
	"GoEncrypt/pkg/cypher"
	"GoEncrypt/pkg/logs"
	"GoEncrypt/pkg/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var CmdDecrypt *cobra.Command

// decyphered a single encrypted file
func RunDecrypt(cmd *cobra.Command, args []string) error {
	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return fmt.Errorf("[!] ERROR : no file path founded")
	}

	fmt.Println("\n------------------------")
	fmt.Printf("Deciphering file %s \n", filePath)

	dir := filepath.Dir(filePath)
	ext := filepath.Ext(filePath) // Extraire l'extension du fichier d'entrée
	outputPath := filepath.Join(dir, fmt.Sprintf("decipher_%s%s", uuid.New().String(), ext))

	rootPath, err := utils.GetRootPath("go.mod")
	if err != nil {
		return fmt.Errorf("[!] ERROR : root file not found")
	}

	privateKeyPath := filepath.Join(rootPath, "data/keys", "private.pem")

	plaintext, err := cypher.DecryptFile(filePath, privateKeyPath)
	if err != nil {
		return fmt.Errorf("[!] ERROR : encrypted file not found")
	}

	// Écrire le contenu déchiffré dans un nouveau fichier
	writeFile := os.WriteFile(outputPath, plaintext, 0644)
	if writeFile != nil {
		return fmt.Errorf("[!] ERROR : no file path founded")
	}

	logs.WriteLogs("[DECRYPT] : decrypt file : " + filePath)

	fmt.Println("\n\n-------------------------")
	fmt.Println("File successullfy decyphered.")
	fmt.Printf("New file : %s", outputPath)
	return nil
}

// decrypt cmd.
func init() {
	CmdDecrypt = &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt a given file using it's path",
		RunE:  RunDecrypt,
	}

	CmdDecrypt.Flags().StringP("file", "f", "", "Path to the file to decrypt")
	CmdDecrypt.MarkFlagRequired("file")

	rootCmd.AddCommand(CmdDecrypt)
}
