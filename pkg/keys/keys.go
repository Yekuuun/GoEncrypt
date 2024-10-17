/*
* Notes : this file contains base code to manage keys in GoEncrypt
 */

package keys

import (
	"GoEncrypt/pkg/utils"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
)

// Generating pair of RSA
func GenerateRSAKeys(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// save private key to file.
func SavePrivateKey(privatekey *rsa.PrivateKey, path string) error {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)
	return os.WriteFile(path, privateKeyPEM, 0600)
}

// save public key to file.
func SavePublicKeyToFile(publicKey *rsa.PublicKey, path string) error {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)
	return os.WriteFile(path, publicKeyPEM, 0644) // Utilisation de os.WriteFile à la place de ioutil.WriteFile
}

// adding keys to file in /data/keys.
func AddKeysToFile(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) error {
	rootPath, err := utils.GetRootPath("go.mod")
	if err != nil {
		return err
	}

	keyFolder := filepath.Join(rootPath, "data/keys")

	errSavePrivate := SavePrivateKey(privateKey, filepath.Join(keyFolder, "private.pem"))
	if errSavePrivate != nil {
		return errSavePrivate
	}

	errSavePublic := SavePublicKeyToFile(publicKey, filepath.Join(keyFolder, "public.pem"))
	if errSavePublic != nil {
		return errSavePublic
	}

	return nil
}

// clean folder containing keys.
func CleanKeys() error {
	rootPath, err := utils.GetRootPath("go.mod")
	if err != nil {
		return err
	}

	keyFolder := filepath.Join(rootPath, "data/keys")

	entries, err := os.ReadDir(keyFolder)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(keyFolder, entry.Name())

		err = os.RemoveAll(entryPath)
		if err != nil {
			return err
		}
	}

	return nil
}
