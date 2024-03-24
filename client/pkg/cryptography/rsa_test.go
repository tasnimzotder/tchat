package cryptography

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyPair(t *testing.T) {
	bits := 2048

	privateKey, publicKey, err := GenerateKeyPair(bits)
	assert.NoError(t, err)

	// Validate private key
	err = privateKey.Validate()
	assert.NoError(t, err)

	// Encrypt and decrypt a message using the keys
	message := []byte("Hello, World!")
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	assert.NoError(t, err)

	decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	assert.NoError(t, err)

	assert.Equal(t, message, decryptedMessage)
}
