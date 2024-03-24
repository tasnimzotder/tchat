package storage

import "github.com/tasnimzotder/tchat/_client/pkg/models"

func (s *SQLiteStorage) SaveMessage(message models.Message) error {
	result := s.db.Create(message)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SQLiteStorage) SaveMessages(messages []models.Message) error {
	if len(messages) == 0 {
		return nil
	}

	result := s.db.Create(messages)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SQLiteStorage) GetMessages() ([]models.Message, error) {
	var messages []models.Message
	result := s.db.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}

func (s *SQLiteStorage) DeleteMessages() error {
	result := s.db.Where("1 = 1").Delete(&models.Message{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}