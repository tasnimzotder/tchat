package storage

import "github.com/tasnimzotder/tchat/_client/pkg/models"

func (s *Storage) SaveUser(user models.User) error {
	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Storage) GetLastUser() (models.User, error) {
	var user models.User
	result := s.db.Last(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (s *Storage) DeleteUser(user models.User) error {
	result := s.db.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Storage) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (s *Storage) GetUserID() (string, error) {
	var user models.User
	result := s.db.Last(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.ID, nil
}

func (s *Storage) IsUserExist() (bool, error) {
	var user models.User
	result := s.db.Last(&user)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
