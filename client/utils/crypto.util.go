package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func EncryptMessage(raw []byte) ([]byte, error) {
	key := []byte("example key 1234")

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Error creating cipher: %v", err)
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("Error creating GCM: %v", err)
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Printf("Error creating nonce: %v", err)
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, raw, nil)

	// convert to string
	//cipherTextString := string(cipherText)

	return cipherText, nil
}

func DecryptMessage(cipherText []byte) ([]byte, error) {
	key := []byte("example key 1234")

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Error creating cipher: %v", err)
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("Error creating GCM: %v", err)
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Printf("Error decrypting message: %v", err)
		return nil, err
	}

	return plainText, nil
}

func EncodeBase64(data []byte) string {
	encodedData := base64.StdEncoding.EncodeToString(data)
	return encodedData
}

func DecodeBase64(data string) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return decodedData, nil
}
