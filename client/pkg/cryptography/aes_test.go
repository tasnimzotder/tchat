package cryptography

import (
	"bytes"
	"testing"
)

func TestAesEncryptMessage(t *testing.T) {
	message := []byte("Hello, World!")
	key := []byte("0123456789abcdef") // 16-byte key

	cipherText, err := aesEncryptMessage(message, key)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Decrypt the cipher text using the same key
	decryptedText, err := aesDecryptMessage(cipherText, key)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !bytes.Equal(message, decryptedText) {
		t.Errorf("Decrypted text does not match original message")
	}
}

func TestAesDecryptMessage(t *testing.T) {
	message := []byte("Hello, World!")
	key := []byte("0123456789abcdef") // 16-byte key

	cipherText, err := aesEncryptMessage(message, key)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	plainText, err := aesDecryptMessage(cipherText, key)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !bytes.Equal(plainText, message) {
		t.Errorf("Decrypted text does not match expected plain text")
	}
}