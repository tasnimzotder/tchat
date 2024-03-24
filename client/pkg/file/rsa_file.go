package file

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func StoreRSAKeys(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) error {
	privatePEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	publicPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	privateFileName := "private.pem"
	publicFileName := "public.pem"

	rootPath := GetProjectStoragePath() + "/rsa_keys"

	// create the directory if it doesn't exist
	if err := os.MkdirAll(rootPath, 0755); err != nil {
		return err
	}

	// private key
	privateFile, err := os.Create(rootPath + "/" + privateFileName)
	if err != nil {
		return err
	}

	defer privateFile.Close()

	if err := pem.Encode(privateFile, privatePEM); err != nil {
		return err
	}

	// public key
	publicFile, err := os.Create(rootPath + "/" + publicFileName)
	if err != nil {
		return err
	}

	defer publicFile.Close()

	if err := pem.Encode(publicFile, publicPEM); err != nil {
		return err
	}

	return nil
}


// func GetPrivate