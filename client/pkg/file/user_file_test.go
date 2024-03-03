package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tasnimzotder/tchat/client/models"
)

func TestGetContactByName(t *testing.T) {
	contacts := []models.Contact{
		{
			ID:   "1",
			Name: "name1",
			Key:  "/path/to/key1.pem",
		},
		{
			ID:   "2",
			Name: "name2",
			Key:  "/path/to/key2.pem",
		},
	}

	// mock function
	ReadFromContactFile = func() ([]models.Contact, error) {
		return contacts, nil
	}

	for _, c := range contacts {
		t.Run(c.Name, func(t *testing.T) {
			contact, err := GetContactByName(c.Name)
			assert.NoError(t, err)
			assert.Equal(t, c, contact)
		})
	}
}

func TestGetContactByID(t *testing.T) {
	contacts := []models.Contact{
		{
			ID:   "1",
			Name: "name1",
			Key:  "/path/to/key1.pem",
		},
		{
			ID:   "2",
			Name: "name2",
			Key:  "/path/to/key2.pem",
		},
	}

	// mock function
	ReadFromContactFile = func() ([]models.Contact, error) {
		return contacts, nil
	}

	for _, c := range contacts {
		t.Run(c.ID, func(t *testing.T) {
			contact, err := GetContactByID(c.ID)
			assert.NoError(t, err)
			assert.Equal(t, c, contact)
		})
	}
}
