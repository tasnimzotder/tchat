package file

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tasnimzotder/tchat/client/models"
)

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

		// write empty file
		_, err = file.Write([]byte("[]"))
		if err != nil {
			log.Printf("Failed to write to file: %v", err)
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
