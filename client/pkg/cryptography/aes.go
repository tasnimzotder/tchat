package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
)

func aesEncryptMessage(message, key []byte) ([]byte, error) {
	//validate key length
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		log.Printf("Invalid key length: %d", len(key))
		return nil, errors.New("invalid key length")
	}

	//create new cipher block
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

	cipherText := gcm.Seal(nonce, nonce, message, nil)

	return cipherText, nil
}

func aesDecryptMessage(cipherText, key []byte) ([]byte, error) {
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
