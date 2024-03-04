package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"log"
)

// Encryptioner is an interface that defines the methods for encryption and decryption.
type Encryptioner interface {
	GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)
	EncryptMessage(message []byte, publicKey *rsa.PublicKey) ([]byte, error)
	DecryptMessage(cipherText []byte, privateKey *rsa.PrivateKey) ([]byte, error)
	EncodeBase64(data []byte) string
	DecodeBase64(data string) ([]byte, error)
}

// RSAEncryption is a struct that implements the Encryptioner interface using RSA encryption.
type RSAEncryption struct{}

// aesEncryptMessage encrypts a message using AES encryption.
//
// It takes in a message and a key as []byte and returns the encrypted message as []byte and an error.
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

// aesDecryptMessage decrypts a message using AES encryption.
//
// It takes in a cipherText and a key as []byte and returns the decrypted message as []byte and an error.
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

// EncodeBase64 encodes a byte array to a base64 string.
//
// It takes in a data as []byte and returns the encoded string.
func (r *RSAEncryption) EncodeBase64(data []byte) string {
	encodedData := base64.StdEncoding.EncodeToString(data)
	return encodedData
}

// DecodeBase64 decodes a base64 string to a byte array.
//
// It takes in a data as string and returns the decoded byte array and an error.
func (r *RSAEncryption) DecodeBase64(data string) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return decodedData, nil
}

// GenerateKeyPair generates a new RSA key pair.
//
// It takes in the number of bits for the key and returns the private key, public key, and an error.
func (r *RSAEncryption) GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	// generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	// validate private key for sanity
	err = privateKey.Validate()
	if err != nil {
		return nil, nil, err
	}

	// generate public key from private key
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

// EncryptMessage encrypts a message using RSA and AES encryption.
//
// It takes in a message as []byte and a public key as *rsa.PublicKey and returns the encrypted message as []byte and an error.
func (r *RSAEncryption) EncryptMessage(message []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	aesKey := make([]byte, 32) // 256-bit key
	if _, err := rand.Read(aesKey); err != nil {
		return nil, err
	}

	cipherText, err := aesEncryptMessage(message, aesKey)
	if err != nil {
		return nil, err
	}

	rng := rand.Reader
	encryptedKey, err := rsa.EncryptOAEP(sha256.New(), rng, publicKey, aesKey, nil)
	if err != nil {
		return nil, err
	}

	return append(encryptedKey, cipherText...), nil
}

// DecryptMessage decrypts a message using RSA and AES decryption.
//
// It takes in a cipherText as []byte and a private key as *rsa.PrivateKey and returns the decrypted message as []byte and an error.
func (r *RSAEncryption) DecryptMessage(cipherText []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	encryptedKeySize := privateKey.Size()
	encryptedKey := cipherText[:encryptedKeySize]
	aesCipherText := cipherText[encryptedKeySize:]

	rng := rand.Reader
	aesKey, err := rsa.DecryptOAEP(sha256.New(), rng, privateKey, encryptedKey, nil)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesDecryptMessage(aesCipherText, aesKey)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
