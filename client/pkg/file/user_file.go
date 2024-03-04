package file

import "github.com/tasnimzotder/tchat/client/models"

func GetContactByName(name string) (models.Contact, error) {
	contacts, err := ReadFromContactFile()
	if err != nil {
		return models.Contact{}, err
	}

	for _, contact := range contacts {
		if contact.Name == name {
			return contact, err
		}
	}

	return models.Contact{}, err
}

func GetContactByID(id string) (models.Contact, error) {
	contacts, err := ReadFromContactFile()
	if err != nil {
		return models.Contact{}, err
	}

	for _, contact := range contacts {
		if contact.ID == id {
			return contact, err
		}
	}

	return models.Contact{}, err
}
