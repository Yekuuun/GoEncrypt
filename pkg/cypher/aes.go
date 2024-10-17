package cypher

import (
	"GoEncrypt/pkg/utils"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"io"
)

func EncryptAES(filePath string, aesKey []byte) ([]byte, []byte, error) {
	plaintext, err := utils.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	return nonce, ciphertext, nil
}

func EncryptAESKeyWithRSA(publicKey *rsa.PublicKey, aesKey []byte) ([]byte, error) {
	encryptedKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, aesKey, nil)
	if err != nil {
		return nil, err
	}
	return encryptedKey, nil
}

func DecryptAES(aesKey, nonce, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// decrypt a single file.
func DecryptFile(encryptedFilePath, privateKeyPath string) ([]byte, error) {
	encryptedData, err := utils.ReadFile(encryptedFilePath)
	if err != nil {
		return nil, err
	}

	rsaPrivateKey, err := utils.LoadRSAPrivateKeyFromPEM(privateKeyPath)
	if err != nil {
		return nil, err
	}

	aesKeySize := rsaPrivateKey.Size()
	encryptedAESKey := encryptedData[:aesKeySize]
	nonce := encryptedData[aesKeySize : aesKeySize+12]
	ciphertext := encryptedData[aesKeySize+12:]

	aesKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaPrivateKey, encryptedAESKey, nil)
	if err != nil {
		return nil, err
	}

	plaintext, err := DecryptAES(aesKey, nonce, ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
