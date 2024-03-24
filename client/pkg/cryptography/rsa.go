package cryptography

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
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

func ConvertRSAToBytes(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) ([]byte, []byte) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)

	return privateKeyBytes, publicKeyBytes
}

func ConvertPrivateBytesToRSA(privateKeyBytes []byte) (*rsa.PrivateKey, error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func ConvertPublicBytesToRSA(publicKeyBytes []byte) (*rsa.PublicKey, error) {
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
