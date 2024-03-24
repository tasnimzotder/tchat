package storage

import "github.com/tasnimzotder/tchat/_client/pkg/models"

// rsa keys
func (s *SQLiteStorage) SaveRSAKeys(privateKey, publicKey []byte) error {
	result := s.db.Create(&models.RSAKeys{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SQLiteStorage) DeleteRSAKeys() error {
	result := s.db.Where("id >= 0").Delete(&models.RSAKeys{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SQLiteStorage) GetPublicRSAKey() ([]byte, error) {
	var rsaKeys models.RSAKeys
	result := s.db.First(&rsaKeys)
	if result.Error != nil {
		return nil, result.Error
	}

	return rsaKeys.PublicKey, nil
}

func (s *SQLiteStorage) GetPrivateRSAKey() ([]byte, error) {
	var rsaKeys models.RSAKeys
	result := s.db.First(&rsaKeys)
	if result.Error != nil {
		return nil, result.Error
	}

	return rsaKeys.PrivateKey, nil
}
