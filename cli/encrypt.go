package cli

import (
	"GoEncrypt/pkg/cypher"
	"GoEncrypt/pkg/utils"
	"crypto/rand"
	"fmt"

	"github.com/spf13/cobra"
)

var CmdEncrypt *cobra.Command

// running file encryption
func RunEncryption(cmd *cobra.Command, args []string) error {
	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return fmt.Errorf("[!] ERROR : no file path founded")
	}

	aesKey := make([]byte, 32)
	if _, err := rand.Read(aesKey); err != nil {
		return fmt.Errorf("[!] ERROR : error on AES key generation")
	}

	nonce, ciphertext, err := cypher.EncryptAES(filePath, aesKey)
	if err != nil {
		return fmt.Errorf("[!] ERROR : error on cyphering")
	}

	publicKey, err := utils.LoadRSAPublicKeyFromPEM("public.pem")
	if err != nil {
		return fmt.Errorf("[!] ERROR : error on loading keys")
	}

	encryptedAES, err := cypher.EncryptAESKeyWithRSA(publicKey, aesKey)
	if err != nil {
		return fmt.Errorf("[!] ERROR : error for aes key encryption")
	}

	outputFilePath, err := utils.BuildEncryptionFileNamePath(filePath)
	if err != nil {
		return fmt.Errorf("[!] ERROR : error building new encrypted file")
	}

	err = utils.SaveEncryptedData(outputFilePath, encryptedAES, nonce, ciphertext)
	if err != nil {
		return fmt.Errorf("[!] ERROR : error uploading new encrypted file")
	}

	fmt.Println("\n------------------------")
	fmt.Println("File succesffuly encrypted")
	return nil
}

func init() {
	CmdEncrypt = &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a given file using it's path",
		RunE:  RunEncryption,
	}

	CmdEncrypt.Flags().StringP("file", "f", "", "Path to the file to encrypt")
	CmdEncrypt.MarkFlagRequired("file")

	rootCmd.AddCommand(CmdEncrypt)
}
