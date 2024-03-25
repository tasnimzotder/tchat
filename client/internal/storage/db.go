package storage

import (
	"github.com/tasnimzotder/tchat/_client/pkg/client"
	"github.com/tasnimzotder/tchat/_client/pkg/file"
	"github.com/tasnimzotder/tchat/_client/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

type Storage struct {
	db  *gorm.DB
	API *client.Client
}

func NewStorage(client *client.Client) (*Storage, error) {
	dbPath := file.GetProjectStoragePath() + "/tchat.db"

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		_, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Storage{
		db:  db,
		API: client,
	}, nil
}

func (s *Storage) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (s *Storage) Migrate() error {
	err := s.db.AutoMigrate(
		&models.User{},
		&models.RSAKeys{},
		&models.Contact{},
		&models.Message{},
	)
	if err != nil {
		return err
	}

	return nil
}
