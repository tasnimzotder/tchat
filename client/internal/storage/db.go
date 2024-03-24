package storage

import (
	"github.com/tasnimzotder/tchat/_client/pkg/file"
	"github.com/tasnimzotder/tchat/_client/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteStorage struct {
	db *gorm.DB
}

func NewSQLiteStorage() (*SQLiteStorage, error) {
	// if dbPath == "" || dbPath[len(dbPath)-7:] != ".sqlite" {
	// 	return nil, errors.New("invalid database file")
	// }

	dbPath := file.GetProjectStoragePath() + "/tchat.db"

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{
		db: db,
	}, nil
}

func (s *SQLiteStorage) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (s *SQLiteStorage) Migrate() error {
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

