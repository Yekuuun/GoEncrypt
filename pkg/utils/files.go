package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// reading a single file
func ReadFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("[!] ERROR : file do not exist")
			return nil, err
		}
		return nil, err
	}

	//reading file
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

// creating encrypted file
func SaveEncryptedData(filePath string, encryptedAESKey, nonce, ciphertext []byte) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(encryptedAESKey)
	f.Write(nonce)
	f.Write(ciphertext)

	return nil
}

// loading rsa key from file
func LoadRSAPublicKeyFromPEM(fileName string) (*rsa.PublicKey, error) {
	rootPath, err := GetRootPath("go.mod")
	if err != nil {
		return nil, err
	}

	keyFolder := filepath.Join(rootPath, "data/keys")

	pubPEM, err := ReadFile(filepath.Join(keyFolder, fileName))
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pubPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("non valid public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not a valid RSA")
	}

	return rsaPubKey, nil
}

// loading private key
func LoadRSAPrivateKeyFromPEM(fileName string) (*rsa.PrivateKey, error) {
	rootPath, err := GetRootPath("go.mod")
	if err != nil {
		return nil, err
	}

	keyFolder := filepath.Join(rootPath, "data/keys")

	privPEM, err := ReadFile(filepath.Join(keyFolder, "private.pem"))
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privPEM)
	if block == nil {
		return nil, err
	}
	if block.Type != "PRIVATE KEY" && block.Type != "RSA PRIVATE KEY" {
		return nil, err
	}

	var privKey *rsa.PrivateKey
	if block.Type == "PRIVATE KEY" {
		parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		var ok bool
		if privKey, ok = parsedKey.(*rsa.PrivateKey); !ok {
			return nil, err
		}
	} else {
		var err error
		privKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}

	return privKey, nil
}

// building path for new encrypted file
func BuildEncryptionFileNamePath(filePath string) (string, error) {
	dir := filepath.Dir(filePath)

	baseName := filepath.Base(filePath)
	fileNameWithoutEx := strings.TrimSuffix(baseName, filepath.Ext(baseName))

	encryptedFileName := fileNameWithoutEx + "_encrypted" + filepath.Ext(filePath)

	encryptedFilePath := filepath.Join(dir, encryptedFileName)

	return encryptedFilePath, nil
}
