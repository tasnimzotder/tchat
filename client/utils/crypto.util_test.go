package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyPair(t *testing.T) {
	bits := []int{1024, 2048}

	for _, bit := range bits {
		t.Run(fmt.Sprintf("%d bits", bit), func(t *testing.T) {
			privateKey, publicKey, err := GenerateKeyPair(bit)

			assert.NoError(t, err)
			assert.NotNil(t, privateKey)
			assert.NotNil(t, publicKey)
		})
	}
}

func TestGenerateKeyPairErrors(t *testing.T) {
	_, _, err := GenerateKeyPair(0)

	assert.Error(t, err)
}

func TestAESEncryptDecrypt(t *testing.T) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	plaintext := []byte("test")

	cipherTest, err := aesEncryptMessage(plaintext, key)
	assert.NoError(t, err)

	decrypted, err := aesDecryptMessage(cipherTest, key)
	assert.NoError(t, err)

	assert.Equal(t, plaintext, decrypted)
}

func TestAESEncryptDecryptErrors(t *testing.T) {
	var keys [][]byte
	plaintext := []byte("test")
	keyLengths := []int{11, 13, 21, 23, 33}

	for _, keyLength := range keyLengths {
		key := make([]byte, keyLength)
		if _, err := rand.Read(key); err != nil {
			log.Fatal(err)
		}

		keys = append(keys, key)
	}

	for _, key := range keys {
		t.Run(fmt.Sprintf("key_%v", len(key)), func(t *testing.T) {
			_, err := aesEncryptMessage(plaintext, key)
			assert.Error(t, err)

			_, err = aesDecryptMessage(plaintext, key)
			assert.Error(t, err)
		})
	}
}

func TestEncryptDecryptMessage(t *testing.T) {
	// process
	// 1. generate key pair
	// 2. encrypt message
	// 3. encode cipher text to base64
	// 4. decode base64 to cipher text
	// 5. decrypt cipher text

	privateKey, publicKey, err := GenerateKeyPair(2048)
	assert.NoError(t, err)

	plaintext := []byte("test")

	// encrypt
	cipherBytes, err := EncryptMessage(plaintext, publicKey)
	assert.NoError(t, err)

	encodedBytes := EncodeBase64(cipherBytes)
	assert.NotEmpty(t, encodedBytes)

	// decrypt
	decodedBytes, err := DecodeBase64(encodedBytes)
	assert.NoError(t, err)
	assert.NotEmpty(t, decodedBytes)

	decrypted, err := DecryptMessage(decodedBytes, privateKey)
	assert.NoError(t, err)

	assert.Equal(t, plaintext, decrypted, "Round trip failed, got %s, want %s", decrypted, plaintext)
}
