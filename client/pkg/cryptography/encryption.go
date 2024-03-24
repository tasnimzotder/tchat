package cryptography

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func EncryptMessage(message []byte, publicKey *rsa.PublicKey) ([]byte, error) {
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

func DecryptMessage(cipherText []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
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

func EncryptMessageWithRSAString(message []byte, publicKeyString string) (string, error) {
	decodedPublicKey, err := DecodeBase64(publicKeyString)
	if err != nil {
		return "", err
	}

	rsaPublicKey, err := ConvertPublicBytesToRSA(decodedPublicKey)
	if err != nil {
		return "", err
	}

	encryptedMessage, err := EncryptMessage(message, rsaPublicKey)
	if err != nil {
		return "", err
	}

	encodedEncryptedMessage := EncodeBase64(encryptedMessage)

	return encodedEncryptedMessage, nil
}
