package storage

import "github.com/tasnimzotder/tchat/_client/pkg/models"

// contacts
func (s *SQLiteStorage) SaveContact(contact models.Contact) error {
	// check if contact and id exists
	var c models.Contact

	// don't throw error if contact doesn't exist
	result := s.db.Where("id = ?", contact.ID).Or("name = ?", contact.Name).First(&c)

	// if contact exists, update it
	if result.Error == nil {
		result = s.db.Model(&c).Updates(contact)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}

	// if contact doesn't exist, create it
	result = s.db.Create(contact)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SQLiteStorage) GetContacts() ([]models.Contact, error) {
	var contacts []models.Contact
	result := s.db.Find(&contacts)
	if result.Error != nil {
		return nil, result.Error
	}

	return contacts, nil
}

func (s *SQLiteStorage) GetContactByName(name string) (models.Contact, error) {
	var contact models.Contact
	result := s.db.Where("name = ?", name).First(&contact)
	if result.Error != nil {
		return models.Contact{}, result.Error
	}

	return contact, nil
}

func (s *SQLiteStorage) GetContactByID(id string) (models.Contact, error) {
	var contact models.Contact
	result := s.db.Where("id = ?", id).First(&contact)
	if result.Error != nil {
		return models.Contact{}, result.Error
	}

	return contact, nil
}
