package cryptography

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptMessage(t *testing.T) {
	message := []byte("Hello, World!")
	_, publicKey, err := GenerateKeyPair(2048)
	assert.NoError(t, err)

	encryptedData, err := EncryptMessage(message, publicKey)

	assert.NoError(t, err)
	assert.NotNil(t, encryptedData)
}


func TestDecryptMessage(t *testing.T) {
	message := []byte("Hello, World!")
	// privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// assert.NoError(t, err)
	privateKey, _, err := GenerateKeyPair(2048)
	assert.NoError(t, err)

	encryptedData, err := EncryptMessage(message, &privateKey.PublicKey)
	assert.NoError(t, err)

	decryptedData, err := DecryptMessage(encryptedData, privateKey)
	assert.NoError(t, err)
	assert.Equal(t, message, decryptedData)
}