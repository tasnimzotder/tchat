package cryptography

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeBase64(t *testing.T) {
	data := []byte("Hello, World!")
	expected := base64.StdEncoding.EncodeToString(data)

	result := EncodeBase64(data)

	assert.Equal(t, expected, result)
}

func TestDecodeBase64(t *testing.T) {
	data := "SGVsbG8sIFdvcmxkIQ==" // Base64 encoded "Hello, World!"
	expected := []byte("Hello, World!")

	result, err := DecodeBase64(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}