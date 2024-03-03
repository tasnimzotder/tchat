package file

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
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

func StoreContactPublicKey(userID string, srcKeyPath string) (string, error) {
	destFileName := getConfigDir() + "/contact_keys/" + userID + ".pem"

	//	copy the file to the target directory
	srcFile, err := os.Open(srcKeyPath)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	//	create the destination file
	destFile, err := os.Create(destFileName)
	if err != nil {
		return "", err
	}

	defer destFile.Close()

	//	copy the file
	_, err = io.Copy(destFile, srcFile)

	if err != nil {
		return "", err
	}

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
