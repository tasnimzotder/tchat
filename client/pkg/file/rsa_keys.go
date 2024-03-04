package file

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/tasnimzotder/tchat/client/pkg/crypto"
)

// StoreRSAKeys RSA keys store
func StoreRSAKeys(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) error {
	privatePEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	publicPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	//	save keys to file
	privateFileName := getConfigDir() + "/keys/private.pem"
	publicFileName := getConfigDir() + "/keys/public.pem"

	//	create the files
	privateFile, err := os.Create(privateFileName)
	if err != nil {
		return err
	}
	defer privateFile.Close()

	publicFile, err := os.Create(publicFileName)
	if err != nil {
		return err
	}
	defer publicFile.Close()

	//	write to files
	err = pem.Encode(privateFile, privatePEM)
	err = pem.Encode(publicFile, publicPEM)

	if err != nil {
		return err
	}

	return nil
}

func StoreContactPublicKey(userID string, publicKey string) (string, error) {
	destFileName := getConfigDir() + "/contact_keys/" + userID + ".pem"

	log.Printf("Public key: %s", publicKey)

	var encryptioner crypto.Encryptioner = &crypto.RSAEncryption{}

	decodedPublicKey, err := encryptioner.DecodeBase64(publicKey)
	if err != nil {
		return "", err
	}

	log.Printf("Decoded public key: %s", decodedPublicKey)

	pemBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: decodedPublicKey,
	}

	pemFile, err := os.Create(destFileName)
	if err != nil {
		return "", err
	}
	defer pemFile.Close()

	err = pem.Encode(pemFile, pemBlock)
	if err != nil {
		return "", err
	}

	log.Printf("Public key stored at: %s", destFileName)

	return destFileName, nil
}

func GetPublicKeyByUserID(userID string) (*rsa.PublicKey, error) {
	contacts, err := ReadFromContactFile()
	if err != nil {
		return nil, err
	}

	keyFileName := getConfigDir() + "/contact_keys/" + userID + ".pem"

	for _, contact := range contacts {
		if contact.ID == userID {
			contents, err := GetFileContents(keyFileName)
			if err != nil {
				return nil, err
			}

			block, _ := pem.Decode(contents)
			publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
			if err != nil {
				return nil, err
			}

			return publicKey, nil
		}
	}

	return nil, nil
}

func GetPrivateKey() (*rsa.PrivateKey, error) {
	privateFileName := getConfigDir() + "/keys/private.pem"

	contents, err := GetFileContents(privateFileName)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(contents)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func GetPublicRSAKey() (*rsa.PublicKey, error) {
	publicFileName := getConfigDir() + "/keys/public.pem"

	contents, err := GetFileContents(publicFileName)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(contents)
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
