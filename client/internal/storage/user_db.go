package storage

import "github.com/tasnimzotder/tchat/_client/pkg/models"

func (s *SQLiteStorage) SaveUser(user models.User) error {
	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SQLiteStorage) GetLastUser() (models.User, error) {
	var user models.User
	result := s.db.Last(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (s *SQLiteStorage) DeleteUser(user models.User) error {
	result := s.db.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SQLiteStorage) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (s *SQLiteStorage) GetUserID() (string, error) {
	var user models.User
	result := s.db.Last(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.ID, nil
}
