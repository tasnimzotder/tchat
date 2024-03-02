package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io"
	"log"
	"os"

	"github.com/tasnimzotder/tchat/client/models"
)

func getConfigDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/tchat"
}

func ClearMessagesFile() error {
	fileName := getConfigDir() + "/messages.json"

	// open file
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// clear the file before writing
	err = file.Truncate(0)
	if err != nil {
		log.Printf("Failed to truncate file: %v", err)
		return err
	}

	// write to file as json empty array
	err = json.NewEncoder(file).Encode([]models.Message{})
	if err != nil {
		log.Printf("Failed to encode messages: %v", err)
		return err
	}

	return nil
}

func ReadFromMessagesFile() ([]models.Message, error) {
	fileName := getConfigDir() + "/messages.json"

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// read from file as json
	var messages []models.Message
	err = json.NewDecoder(file).Decode(&messages)
	if err != nil {
		//log.Printf("Failed to decode messages: %v", err)
		return nil, err
	}

	return messages, nil
}

func AppendToMessagesFile(message models.Message) error {
	fileName := getConfigDir() + "/messages.json"

	// create if not exists
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		// create the directory

		file, err := os.Create(fileName)
		if err != nil {
			return err
		}

		defer file.Close()
	}

	messages, err := ReadFromMessagesFile()
	if err != nil {
		log.Printf("Failed to read from messages file: %v", err)
		return err
	}

	messages = append(messages, message)

	// open file
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// clear the file before writing
	err = file.Truncate(0)
	if err != nil {
		log.Printf("Failed to truncate file: %v", err)
		return err
	}

	// write to file as json
	err = json.NewEncoder(file).Encode(messages)
	if err != nil {
		log.Printf("Failed to encode messages: %v", err)
		return err
	}

	return nil
}

func WriteToConfigFile(config models.Config) error {
	fileName := getConfigDir() + "/config.json"

	// create if not exists
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		// create the directory

		file, err := os.Create(fileName)
		if err != nil {
			return err
		}

		defer file.Close()
	}
	// open file
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	// clear the file before writing
	err = file.Truncate(0)
	if err != nil {
		log.Printf("Failed to truncate file: %v", err)
		return err
	}

	// write to file as json
	err = json.NewEncoder(file).Encode(config)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromConfigFile() (models.Config, error) {
	fileName := getConfigDir() + "/config.json"

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		return models.Config{}, err
	}

	defer file.Close()

	// read from file as json
	var config models.Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return models.Config{}, err
	}

	return config, nil
}

func GetFileContents(fileName string) ([]byte, error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return nil, err
	}

	return contents, nil
}

func WriteToContactFile(contact models.Contact) error {
	fileName := getConfigDir() + "/contacts.json"

	// create if not exists
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		// create the directory

		file, err := os.Create(fileName)
		if err != nil {
			log.Printf("Failed to create file: %v", err)
			return err
		}

		defer file.Close()
	}

	// open file
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return err
	}

	defer file.Close()

	contacts, err := ReadFromContactFile()
	if err != nil {
		log.Printf("Failed to read from contacts file: %v", err)
		return err

	}

	exist := false

	for i, c := range contacts {
		if c.ID == contact.ID {
			contacts[i] = contact
			exist = true
			break
		}
	}

	if !exist {
		contacts = append(contacts, contact)
	}

	// clear the file before writing
	err = file.Truncate(0)
	if err != nil {
		log.Printf("Failed to truncate file: %v", err)
		return err
	}

	// write to file as json
	err = json.NewEncoder(file).Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

var ReadFromContactFile = func() ([]models.Contact, error) {
	//func ReadFromContactFile() ([]models.Contact, error) {
	fileName := getConfigDir() + "/contacts.json"

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// read from file as json
	var contacts []models.Contact
	err = json.NewDecoder(file).Decode(&contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

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
